package xdb

import "fmt"

func CreateSearcher(dbPath string, cachePolicy string) (*Searcher, error) {
	switch cachePolicy {
	case "nil", "file":
		return NewWithFileOnly(dbPath)
	case "vectorIndex":
		vIndex, err := LoadVectorIndexFromFile(dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load vector index from `%s`: %w", dbPath, err)
		}

		return NewWithVectorIndex(dbPath, vIndex)
	case "content":
		cBuff, err := LoadContentFromFile(dbPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load content from '%s': %w", dbPath, err)
		}

		return NewWithBuffer(cBuff)
	default:
		return nil, fmt.Errorf("invalid cache policy `%s`, options: file/vectorIndex/content", cachePolicy)
	}
}
