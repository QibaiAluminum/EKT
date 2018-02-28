package MPTPlus

import (
	"bytes"
	"sort"

	"github.com/EducationEKT/EKT/io/ekt8/db"
)

/*
*TrieSonInfo存储的是当前节点的下一个节点的信息
*如果当前节点是叶子节点,则Sons的长度为1,且TrieSonInfo.Hash是当前路径key的Value值的hash
*如果当前节点不是叶子节点,则Sons的长度大于等于1,存储的是子节点的Hash值和PathValue
 */
type TrieSonInfo struct {
	Hash      []byte
	PathValue []byte
}

/*
*TrieNode存储的是当前节点的一些信息,包括PathValue 是否是叶子节点,子节点信息等等
*strings.Join(pathValue,"")就是用户要存储的key
 */
type TrieNode struct {
	Sons      SortedSon
	Leaf      bool
	Root      bool
	PathValue []byte
}

type MTP struct {
	Root []byte
	DB   db.EKTDB
}

type SortedSon []TrieSonInfo

func (sonInfo SortedSon) Len() int {
	return len(sonInfo)
}

func (sonInfo SortedSon) Swap(i, j int) {
	sonInfo[i], sonInfo[j] = sonInfo[j], sonInfo[i]
}

func (sonInfo SortedSon) Less(i, j int) bool {
	length := len(sonInfo[i].PathValue)
	if len(sonInfo[j].PathValue) < length {
		length = len(sonInfo[j].PathValue)
	}
	for i := 0; i < length; i++ {
		if sonInfo[i].PathValue[i] < sonInfo[j].PathValue[i] {
			return true
		} else if sonInfo[i].PathValue[i] > sonInfo[j].PathValue[i] {
			return false
		}
	}
	return true
}

func (node *TrieNode) AddSon(hash, pathValue []byte) {
	if nil == node.Sons {
		node.Sons = *new(SortedSon)
	}
	for _, son := range node.Sons {
		if bytes.Equal(son.PathValue, pathValue) {
			node.DeleteSon(pathValue)
		}
	}
	node.Sons = append(node.Sons, TrieSonInfo{Hash: hash, PathValue: pathValue})
	sort.Sort(node.Sons)
}

func (node *TrieNode) DeleteSon(pathValue []byte) {
	for i, son := range node.Sons {
		if bytes.Equal(son.PathValue, pathValue) {
			node.Sons = append(node.Sons[:i-1], node.Sons[i:]...)
		}
	}
}