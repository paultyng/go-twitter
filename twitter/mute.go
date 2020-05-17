package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Mutes is a cursored collection of muted users.
type Mutes struct {
	Users             []User `json:"users"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int64  `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
}

// MuteService provides methods for mute specific users.
type MuteService struct {
	sling *sling.Sling
}

// newMuteService returns a new MuteService.
func newMuteService(sling *sling.Sling) *MuteService {
	return &MuteService{
		sling: sling.Path("mutes/users/"),
	}
}

// MuteListParams are the parameters for MuteService.List
type MuteListParams struct {
	Cursor              int64 `url:"cursor,omitempty"`
	SkipStatus          *bool `url:"skip_status,omitempty"`
	IncludeUserEntities *bool `url:"include_user_entities,omitempty"`
}

// List returns a cursored collection of Users blocked by the current user.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/get-mutes-users-list
func (s *MuteService) List(params *MuteListParams) (*Mutes, *http.Response, error) {
	mutes := new(Mutes)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("list.json").QueryStruct(params).Receive(mutes, apiError)
	return mutes, resp, relevantError(err, *apiError)
}

// MuteCreateParams are the parameters for MuteService.Create.
type MuteCreateParams struct {
	ScreenName string `url:"screen_name,omitempty,comma"`
	UserID     int64  `url:"user_id,omitempty,comma"`
}

// Create a mute for specific user, return the user muted as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/post-mutes-users-create
func (s *MuteService) Create(params *MuteCreateParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}

// MuteDestroyParams are the parameters for MuteService.Destroy.
type MuteDestroyParams struct {
	ScreenName string `url:"screen_name,omitempty,comma"`
	UserID     int64  `url:"user_id,omitempty,comma"`
}

// Destroy the mute for specific user, return the user unmuted as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/post-mutes-users-destroy
func (s *MuteService) Destroy(params *MuteDestroyParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}
