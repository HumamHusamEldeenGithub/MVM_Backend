package service

import "fmt"

func (s *mvmService) CheckRoomAvailability(roomId, userId string) error {
	room, err := s.store.GetRoom(roomId)
	if err != nil {
		return err
	}
	if room.OwnerId == userId || (!room.IsPrivate && !room.FriendsOnly) {
		return nil
	}

	invitation := CheckInvitations(userId, room.Invitations)
	if invitation {
		return nil
	}

	if room.FriendsOnly {
		isFriends, _ := s.CheckFriendship(room.OwnerId, userId)
		if isFriends {
			return nil
		}
	}

	return fmt.Errorf("not authorized to enter this room ")
}

func (s *mvmService) CheckFriendship(user1, user2 string) (bool, error) {
	friends, err := s.store.GetFriends(user1)
	if err != nil {
		return false, err
	}
	for _, friendId := range friends.Friends {
		if friendId == user2 {
			return true, nil
		}
	}
	return false, nil
}

func CheckInvitations(userId string, invitations []string) bool {
	for _, invitation := range invitations {
		if invitation == userId {
			return true
		}
	}
	return false
}
