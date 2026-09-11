package replicator

import "github.com/tamzrod/MCS.OSJS/mma2composer"

func (s Store) loadEffectiveConfig() (mma2composer.EffectiveConfig, error) {
	return mma2composer.New(s.Root, ProducerReplicator).LoadEffective()
}
