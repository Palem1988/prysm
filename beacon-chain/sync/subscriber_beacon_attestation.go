package sync

import (
	"context"

	"github.com/gogo/protobuf/proto"
	ethpb "github.com/prysmaticlabs/prysm/proto/eth/v1alpha1"
)

// beaconAttestationSubscriber forwards the incoming validated attestation to the blockchain
// service for processing.
func (r *RegularSync) beaconAttestationSubscriber(ctx context.Context, msg proto.Message) error {
	return r.chain.ReceiveAttestationNoPubsub(ctx, msg.(*ethpb.Attestation))
}
