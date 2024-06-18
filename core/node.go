package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Node struct {
	Address    string
	Blockchain *Blockchain
	Peers      map[string]*Peer
	Mutex      sync.Mutex
}

type Peer struct {
	Address string
	Conn    net.Conn
}

func NewNode(address string, blockchain *Blockchain) *Node {
	return &Node{
		Address:    address,
		Blockchain: blockchain,
		Peers:      make(map[string]*Peer),
	}
}

func (node *Node) Start() {
	ln, err := net.Listen("tcp", node.Address)
	if err != nil {
		fmt.Printf("Error starting node: %v\n", err)
		return
	}
	defer ln.Close()

	fmt.Printf("Node started at %s\n", node.Address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go node.handleConnection(conn)
	}
}

func (node *Node) handleConnection(conn net.Conn) {
	defer conn.Close()
	peerAddress := conn.RemoteAddr().String()
	peer := &Peer{Address: peerAddress, Conn: conn}
	node.addPeer(peer)

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading message from %s: %v\n", peerAddress, err)
			node.removePeer(peerAddress)
			return
		}
		node.handleMessage(strings.TrimSpace(message), peer)
	}
}

func (node *Node) addPeer(peer *Peer) {
	node.Mutex.Lock()
	defer node.Mutex.Unlock()
	node.Peers[peer.Address] = peer
	fmt.Printf("Peer %s added\n", peer.Address)
}

func (node *Node) removePeer(address string) {
	node.Mutex.Lock()
	defer node.Mutex.Unlock()
	delete(node.Peers, address)
	fmt.Printf("Peer %s removed\n", address)
}

func (node *Node) handleMessage(message string, peer *Peer) {
	fmt.Printf("Message from %s: %s\n", peer.Address, message)
	msg, err := DeserializeMessage(message)
	if err != nil {
		fmt.Printf("Error deserializing message from %s: %v\n", peer.Address, err)
		return
	}

	switch msg.Type {
	case MessageTypeTransaction:
		node.handleTransactionMessage(msg.Payload)
	case MessageTypeBlock:
		node.handleBlockMessage(msg.Payload)
	default:
		fmt.Printf("Unknown message type from %s: %s\n", peer.Address, msg.Type)
	}
}

func (node *Node) handleTransactionMessage(payload string) {
	var tx Transaction
	err := json.Unmarshal([]byte(payload), &tx)
	if err != nil {
		fmt.Printf("Error unmarshalling transaction: %v\n", err)
		return
	}
	node.Blockchain.AddTransaction(&tx)
	node.broadcastMessage(NewMessage(MessageTypeTransaction, payload).Serialize())
}

func (node *Node) handleBlockMessage(payload string) {
	var block Block
	err := json.Unmarshal([]byte(payload), &block)
	if err != nil {
		fmt.Printf("Error unmarshalling block: %v\n", err)
		return
	}
	node.Blockchain.AddBlock(block.Transactions)
	node.broadcastMessage(NewMessage(MessageTypeBlock, payload).Serialize())
}

func (node *Node) sendMessageToPeer(message, address string) {
	node.Mutex.Lock()
	defer node.Mutex.Unlock()
	if peer, ok := node.Peers[address]; ok {
		_, err := peer.Conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Printf("Error sending message to %s: %v\n", address, err)
		}
	}
}

func (node *Node) broadcastMessage(message string) {
	node.Mutex.Lock()
	defer node.Mutex.Unlock()
	for _, peer := range node.Peers {
		_, err := peer.Conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Printf("Error broadcasting message to %s: %v\n", peer.Address, err)
		}
	}
}
