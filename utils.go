package cfn_custom_resource_helper

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func NewPhysicalResourceID(event CrHelperEvent) string {
	rns := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	snm := strings.Split(event.StackID, "/")[1]
	lid := event.LogicalResourceID
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	rnd := make([]byte, 12)
	for i := range rnd {
		rnd[i] = rns[gen.Intn(len(rns))]
	}
	return fmt.Sprintf("%s-%s-%s", snm, lid, rnd)
}
