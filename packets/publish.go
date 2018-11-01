package packets

import (
	"bytes"
	"io"
	"io/ioutil"
	"net"
)

// Publish is the Variable Header definition for a publish control packet
type Publish struct {
	Duplicate  bool
	QoS        byte
	Retain     bool
	Topic      string
	PacketID   uint16
	Properties *Properties
	Payload    []byte
}

//Unpack is the implementation of the interface required function for a packet
func (p *Publish) Unpack(r *bytes.Buffer) error {
	var err error
	p.Topic, err = readString(r)
	if err != nil {
		return err
	}
	if p.QoS > 0 {
		p.PacketID, err = readUint16(r)
		if err != nil {
			return err
		}
	}

	err = p.Properties.Unpack(r, PUBLISH)
	if err != nil {
		return err
	}

	p.Payload, err = ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return nil
}

// Buffers is the implementation of the interface required function for a packet
func (p *Publish) Buffers() net.Buffers {
	var b bytes.Buffer
	writeString(p.Topic, &b)
	if p.QoS > 0 {
		writeUint16(p.PacketID, &b)
	}
	properties := p.Properties.Pack(PUBLISH)
	propLen := encodeVBI(len(properties))
	return net.Buffers{b.Bytes(), propLen, properties, p.Payload}

}

// WriteTo is the implementation of the interface required function for a packet
func (p *Publish) WriteTo(w io.Writer) (int64, error) {
	f := p.QoS << 1
	if p.Duplicate {
		f |= 1 << 3
	}
	if p.Retain {
		f |= 1
	}

	cp := &ControlPacket{FixedHeader: FixedHeader{Type: PUBLISH, Flags: f}}
	cp.Content = p

	return cp.WriteTo(w)
}
