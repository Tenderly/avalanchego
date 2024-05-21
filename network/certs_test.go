// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"crypto/tls"
	"net"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	_ "embed"

	"github.com/tenderly/avalanchego/ids"
	"github.com/tenderly/avalanchego/network/peer"
	"github.com/tenderly/avalanchego/staking"
	"github.com/tenderly/avalanchego/utils/ips"
)

var (
	//go:embed test_cert_1.crt
	testCertBytes1 []byte
	//go:embed test_key_1.key
	testKeyBytes1 []byte
	//go:embed test_cert_2.crt
	testCertBytes2 []byte
	//go:embed test_key_2.key
	testKeyBytes2 []byte
	//go:embed test_cert_3.crt
	testCertBytes3 []byte
	//go:embed test_key_3.key
	testKeyBytes3 []byte

	ip      *ips.ClaimedIPPort
	otherIP *ips.ClaimedIPPort

	certLock   sync.Mutex
	tlsCerts   []*tls.Certificate
	tlsConfigs []*tls.Config
)

func init() {
	cert1, err := staking.LoadTLSCertFromBytes(testKeyBytes1, testCertBytes1)
	if err != nil {
		panic(err)
	}
	cert2, err := staking.LoadTLSCertFromBytes(testKeyBytes2, testCertBytes2)
	if err != nil {
		panic(err)
	}
	cert3, err := staking.LoadTLSCertFromBytes(testKeyBytes3, testCertBytes3)
	if err != nil {
		panic(err)
	}
	tlsCerts = []*tls.Certificate{
		cert1, cert2, cert3,
	}

	stakingCert1, err := staking.ParseCertificate(cert1.Leaf.Raw)
	if err != nil {
		panic(err)
	}
	stakingCert2, err := staking.ParseCertificate(cert2.Leaf.Raw)
	if err != nil {
		panic(err)
	}

	ip = ips.NewClaimedIPPort(
		stakingCert1,
		ips.IPPort{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 9651,
		},
		1,   // timestamp
		nil, // signature
	)
	otherIP = ips.NewClaimedIPPort(
		stakingCert2,
		ips.IPPort{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 9651,
		},
		1,   // timestamp
		nil, // signature
	)
}

func getTLS(t *testing.T, index int) (ids.NodeID, *tls.Certificate, *tls.Config) {
	certLock.Lock()
	defer certLock.Unlock()

	for len(tlsCerts) <= index {
		cert, err := staking.NewTLSCert()
		require.NoError(t, err)
		tlsCerts = append(tlsCerts, cert)
	}
	for len(tlsConfigs) <= index {
		cert := tlsCerts[len(tlsConfigs)]
		tlsConfig := peer.TLSConfig(*cert, nil)
		tlsConfigs = append(tlsConfigs, tlsConfig)
	}

	tlsCert := tlsCerts[index]
	cert, err := staking.ParseCertificate(tlsCert.Leaf.Raw)
	require.NoError(t, err)
	nodeID := ids.NodeIDFromCert(cert)
	return nodeID, tlsCert, tlsConfigs[index]
}
