package serializers

import "snapchat-clone/models"

func LoginSerializer(user *models.User) *models.User {
	var serialized = models.User{
		Email:        user.Email,
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
	return &serialized
}

func UserDetailSerializer(user *models.User) *models.User {
	var serialized = models.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Birthday:  user.Birthday,
		CreatedAt: user.CreatedAt,
		Timestamp: user.Timestamp,
	}
	return &serialized
}

func ProfileDetailSerializer(profile *models.Profile, user *models.User) *models.Profile {
	user = UserDetailSerializer(user)
	var serialized = models.Profile{
		ID:                    profile.ID,
		User:                  user,
		ProfileImage:          profile.ProfileImage,
		BackgroundImage:       profile.BackgroundImage,
		GhostMode:             profile.GhostMode,
		SeeLocation:           profile.SeeLocation,
		LocationALlFriends:    profile.LocationALlFriends,
		LocationExceptFriends: profile.LocationExceptFriends,
		TwoFactor:             profile.TwoFactor,
		Timestamp:             profile.Timestamp,
	}
	return &serialized
}
