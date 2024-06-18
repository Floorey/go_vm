package tests

import (
	"net"
	"testing"
	"time"
	"vm/core"
)

func TestNodeCommunication(t *testing.T) {
	vm := core.NewVM()
	validatorManager := core.NewValidatorManager()
	poh := core.NewProofOfHistory()

	validatorManager.AddValidator("Validator1", 100)
	validatorManager.AddValidator("Validator2", 50)

	bc := core.NewBlockchain(vm, validatorManager, poh)

	node1 := core.NewNode("localhost:3000", bc)
	go node1.Start()
	time.Sleep(1 * time.Second) // Kurze Verzögerung, um sicherzustellen, dass der Server startet

	node2 := core.NewNode("localhost:3001", bc)
	go node2.Start()
	time.Sleep(1 * time.Second) // Kurze Verzögerung, um sicherzustellen, dass der Server startet

	// Verbindungsaufbau zwischen den Knoten simulieren
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		t.Fatalf("Fehler beim Verbinden zu node1: %v", err)
	}
	defer conn.Close()
	node2.handleConnection(conn)

	// Simulieren Sie das Senden von Nachrichten zwischen den Knoten
	message := core.NewMessage(core.MessageTypeTransaction, `{"Data":"Transfer 100 Coins from Alice to Bob","Code":"TRANSFER"}`).Serialize()
	node1.sendMessageToPeer(message, "localhost:3001")

	// Fügen Sie eine Transaktion hinzu und prüfen Sie, ob sie auf beiden Knoten vorhanden ist
	tx1 := core.NewTransaction("Transfer 100 Coins from Alice to Bob", string([]byte{byte(core.OP_TRANSFER)}))
	node1.Blockchain.AddTransaction(tx1)

	// Warten Sie kurz, damit die Nachricht verarbeitet wird
	time.Sleep(1 * time.Second)

	// Überprüfen Sie die Blockchain beider Knoten
	if len(node1.Blockchain.Blocks[0].Transactions) != 1 {
		t.Fatalf("Erwartete 1 Transaktion in node1, erhielt %d", len(node1.Blockchain.Blocks[0].Transactions))
	}
	if len(node2.Blockchain.Blocks[0].Transactions) != 1 {
		t.Fatalf("Erwartete 1 Transaktion in node2, erhielt %d", len(node2.Blockchain.Blocks[0].Transactions))
	}

	// Print the blockchain (for test verification purposes)
	t.Log("Blockchain:")
	node1.Blockchain.PrintChain()
}
