package serializers

import "snapchat-clone/models"

func LoginSerializer(user *models.User) *models.User {
	var serialized = models.User{
		ID:           user.ID,
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
		UserID:                user.ID,
		ProfileImageURL:       profile.ProfileImageURL,
		BackgroundImageURL:    profile.BackgroundImageURL,
		GhostMode:             profile.GhostMode,
		SeeLocation:           profile.SeeLocation,
		ALlFriends:            profile.ALlFriends,
		LocationExceptFriends: profile.LocationExceptFriends,
		TwoFactor:             profile.TwoFactor,
		Timestamp:             profile.Timestamp,
	}
	return &serialized
}
