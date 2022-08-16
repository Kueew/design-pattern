package prototype

import (
	"time"
)

type SearchWord struct {
	KeyWord  string
	Count    int
	UpdateAt *time.Time
}

func (word *SearchWord) Clone() *SearchWord {
	return &SearchWord{
		KeyWord:  word.KeyWord,
		Count:    word.Count,
		UpdateAt: word.UpdateAt,
	}
}

type SearchWords map[string]*SearchWord

// Clone 拷贝一个新的 SearchWords
// updateWords 需要更新的关键词列表
func (words SearchWords) Clone(updateWords []*SearchWord) SearchWords {
	var newWords = SearchWords{}
	//浅拷贝
	for k, v := range words {
		newWords[k] = v
	}

	for _, v := range updateWords {
		newWords[v.KeyWord] = v.Clone()
	}
	return newWords
}
