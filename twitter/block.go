package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Blocks is a cursored collection of followers.
type Blocks struct {
	Users             []User `json:"users"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int64  `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
}

// BlockService provides methods for blocking specific user.
type BlockService struct {
	sling *sling.Sling
}

// newBlockService returns a new BlockService.
func newBlockService(sling *sling.Sling) *BlockService {
	return &BlockService{
		sling: sling.Path("blocks/"),
	}
}

// BlockListParams are the parameters for BlockService.List
type BlockListParams struct {
	Cursor              int64 `url:"cursor,omitempty"`
	SkipStatus          *bool `url:"skip_status,omitempty"`
	IncludeUserEntities *bool `url:"include_user_entities,omitempty"`
}

// List returns a cursored collection of Users blocked by the current user.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/get-blocks-list
func (s *BlockService) List(params *BlockListParams) (*Blocks, *http.Response, error) {
	blocks := new(Blocks)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("list.json").QueryStruct(params).Receive(blocks, apiError)
	return blocks, resp, relevantError(err, *apiError)
}

// BlockCreateParams are the parameters for BlockService.Create.
type BlockCreateParams struct {
	ScreenName      string `url:"screen_name,omitempty,comma"`
	UserID          int64  `url:"user_id,omitempty,comma"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

// Create a block for specific user, return the user blocked as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/post-blocks-create
func (s *BlockService) Create(params *BlockCreateParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("create.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}

// BlockDestroyParams are the parameters for BlockService.Destroy.
type BlockDestroyParams struct {
	ScreenName      string `url:"screen_name,omitempty,comma"`
	UserID          int64  `url:"user_id,omitempty,comma"`
	IncludeEntities *bool  `url:"include_entities,omitempty"` // whether 'status' should include entities
	SkipStatus      *bool  `url:"skip_status,omitempty"`
}

// Destroy the block for specific user, return the user unblocked as Entity.
// https://developer.twitter.com/en/docs/accounts-and-users/mute-block-report-users/api-reference/post-blocks-destroy
func (s *BlockService) Destroy(params *BlockDestroyParams) (User, *http.Response, error) {
	users := new(User)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("destroy.json").QueryStruct(params).Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}
