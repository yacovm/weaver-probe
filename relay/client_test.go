package relay_test

import (
	"fmt"
	"github.com/yacovm/weaver-probe/relay"
	"testing"
)

func TestClient(t *testing.T) {
	client := relay.Client{}
	res, err := client.RequestState( "127.0.0.1:20040")
	fmt.Println(res, err)
}
