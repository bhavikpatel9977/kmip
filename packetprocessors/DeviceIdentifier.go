package packetprocessors

import "fmt"

type DeviceIdentifier struct {}



func init() {
	Kmpiprocessor["4200A2"] = new(DeviceIdentifier)
}


func (r * DeviceIdentifier) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("DeviceIdentifier",string(t.Tag), string(t.Type) , t.Length, t.Value)

	f,s := ReadTTLV(req)
	p := GetProcessor(string(s.Tag))

	if(len(req[f:])) <= 0 {
		return nil, nil
	}

	if p!= nil {
		p.ProcessPacket(&s,req[f:], nil, nil)
	}
	return nil,nil
}