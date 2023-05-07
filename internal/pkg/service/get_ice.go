package service

import "fmt"

func (s *mvmService) GetIce(id string) ([]string, error) {
	fmt.Println(Clients[id].ICECandidates)
	ices := Clients[id].ICECandidates
	if ices != nil {
		return Clients[id].ICECandidates, nil
	}
	return nil, nil
}
