/*
	NautilusDB is database that has been created to understand how databases work under the hood
    Copyright (C) 2025  Snigdhadeb Roy Chowdhury

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

	For communications or inquiries, contact:
	Email: snigdhadeb_roychowdhury@outlook.com
*/

package index

import "encoding/binary"

const (
	BPLUSNODE_INTERNAL_NODE = 0
	BPLUSNODE_LEAF_NODE     = 1

	HEADER = 4

	BPLUSTREE_PAGE_SIZE      = 4096
	BPLUSTREE_MAX_KEY_SIZE   = 1000
	BPLUSTREE_MAX_VALUE_SIZE = 3000
)

type BPlusNode struct {
	data []byte
}

type BPlusTree struct {
	root uint64                 // pointer to the root/parent node
	get  func(uint64) BPlusNode // dereference a pointer to return the node at that address
	new  func(BPlusNode) uint64 // returns the address of the new node
	del  func(uint64)           // deallocate a page

}

// The size of a node can be HEADER + 8 + 2 + 4 + BPLUSTREE_MAX_KEY_SIZE + BPLUSTREE_MAX_VALUE_SIZE
// It's limit is the BPLUSTREE_PAGE_SIZE

// Below we declare the methods for the BPlusNode type
func (node BPlusNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node.data)
}

func (node BPlusNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node.data[0:2])
}

func (node BPlusNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node.data[0:2], btype)
	binary.LittleEndian.PutUint16(node.data[2:4], nkeys)
}

func (node BPlusNode) getPointer(idx uint16) uint64 {
	pos := HEADER + (8 * idx)
	return binary.LittleEndian.Uint64(node.data[pos:])
}

func (node BPlusNode) setPointer(idx uint16, val uint64) {
	pos := HEADER + (8 * idx)
	binary.LittleEndian.PutUint64(node.data[pos:], val)

}

func offsetPos(node BPlusNode, idx uint16) uint16 {
	return HEADER + (8 * node.nkeys()) + (2 * (idx - 1))
}

func (node BPlusNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}

	return binary.LittleEndian.Uint16(node.data[offsetPos(node, idx):])
}

func (node BPlusNode) setOffset(idx uint16, offset uint16) {
	binary.LittleEndian.PutUint16(node.data[offsetPos(node, idx):], offset)
}

// The offset list is used to locate the nth Key-Value pair quickly

func (node BPlusNode) kvPos(idx uint16) uint16 {
	return HEADER + (8 * node.nkeys()) + (2 * node.nkeys()) + node.getOffset(idx)
}

func (node BPlusNode) getKey(idx uint16) []byte {
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node.data[pos:])
	return node.data[pos+4:][:klen]
}

func (node BPlusNode) getVal(idx uint16) []byte {
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node.data[pos+0:])
	vlen := binary.LittleEndian.Uint16(node.data[pos+2:])
	return node.data[pos+4+klen:][:vlen]
}

func (node BPlusNode) nbytes() uint16 {
	return node.kvPos(node.nkeys())
}
