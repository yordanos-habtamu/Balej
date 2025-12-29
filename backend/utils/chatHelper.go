package utils
import (
	"fmt"
	"sort"
	
)

// GenerateChatID takes two user IDs and returns a consistent ChatID
func GenerateChatID(user1ID, user2ID int) string {
	ids := []int{user1ID, user2ID}
	sort.Ints(ids) // ensures lowest always comes first
	return fmt.Sprintf("%d_%d", ids[0], ids[1])
}
