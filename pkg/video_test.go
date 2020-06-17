package pkg

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestAPIGet(t *testing.T) {
	ServerHost = os.Getenv(`SERVER_HOST`)
	res := APIGetRecorded(1000, 0)

	_, found := res.Videos.First()
	if !found {
		log.Fatal(`not found`)
	}

	videos := res.Videos.Where(func(x Video) bool { return strings.Contains(x.Description, `サッカー`) })
	videos2 := videos.SortBy(func(a Video, b Video) bool { return a.Name > b.Name })
	if videos[0].ID == videos2[0].ID {
		log.Fatalf(`before:%d, after:%d`, videos[0].ID, videos2[0].ID)
	}

	videos3 := videos.Mapper(func(x Video) Video { x.Name = "aaa"; return x })
	if videos3[0].Name != `aaa` {
		log.Fatalf(`Name shoud be %s`, videos3[0].Name)
	}

}
