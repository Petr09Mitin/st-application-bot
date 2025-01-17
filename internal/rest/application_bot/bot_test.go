package applicationbot

// import (
// 	"testing"

// 	"github.com/NicoNex/echotron/v3"
// 	"github.com/golang/mock/gomock"
// 	"github.com/st-application-bot/domain"
// 	"github.com/st-application-bot/internal/config"
// 	mock_applicationbot "github.com/st-application-bot/mocks/rest/application_bot"
// 	"github.com/stretchr/testify/assert"
// )

// type fields struct {
// 	as  *mock_applicationbot.MockApplicationService
// 	api *mock_applicationbot.MockBotAPI
// 	l   *mock_applicationbot.MockLogger
// }

// func TestApplicationBot_Update(t *testing.T) {
// 	type args struct {
// 		u *echotron.Update
// 	}
// 	tests := []struct {
// 		name string
// 		prep func(f *fields)
// 		args args
// 	}{
// 		{
// 			name: "nil update",
// 			prep: func(f *fields) {
// 				f.l.EXPECT().Error(domain.ErrInvalidUpdate.Error()).Times(1)
// 			},
// 			args: args{
// 				u: nil,
// 			},
// 		},
// 		{
// 			name: "empty update",
// 			prep: func(f *fields) {
// 				f.l.EXPECT().Error(domain.ErrInvalidUpdate.Error()).Times(1)
// 			},
// 			args: args{
// 				u: &echotron.Update{},
// 			},
// 		},
// 		{
// 			name: "update with invalid update.Message",
// 			prep: func(f *fields) {
// 				f.l.EXPECT().Error(gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				u: &echotron.Update{
// 					Message: &echotron.Message{},
// 				},
// 			},
// 		},
// 		{
// 			name: "update with invalid update.CallbackQuery",
// 			prep: func(f *fields) {
// 				f.l.EXPECT().Error(gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				u: &echotron.Update{
// 					CallbackQuery: &echotron.CallbackQuery{},
// 				},
// 			},
// 		},
// 		{
// 			name: "update with valid update.Message",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(PressStartText, int64(1), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, nil)
// 			},
// 			args: args{
// 				u: &echotron.Update{
// 					Message: &echotron.Message{
// 						Chat: echotron.Chat{
// 							ID:   1,
// 							Type: TgPrivateChatType,
// 						},
// 						From: &echotron.User{
// 							Username: "test",
// 						},
// 						Text: "test",
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name: "error sending responseMessage",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(PressStartText, int64(1), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, domain.ErrInternal)
// 				f.l.EXPECT().Error(domain.ErrInternal.Error()).Times(1)
// 			},
// 			args: args{
// 				u: &echotron.Update{
// 					Message: &echotron.Message{
// 						Chat: echotron.Chat{
// 							ID:   1,
// 							Type: TgPrivateChatType,
// 						},
// 						From: &echotron.User{
// 							Username: "test",
// 						},
// 						Text: "test",
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			f := &fields{
// 				nil,
// 				mock_applicationbot.NewMockBotAPI(ctrl),
// 				mock_applicationbot.NewMockLogger(ctrl),
// 			}
// 			if tt.prep != nil {
// 				tt.prep(f)
// 			}

// 			b := &ApplicationBot{
// 				1,
// 				1,
// 				f.as,
// 				&domain.Application{},
// 				InitialState,
// 				f.l,
// 				nil,
// 				f.api,
// 			}
// 			b.Update(tt.args.u)
// 		})
// 	}
// }

// func TestApplicationBot_setSenderInfo(t *testing.T) {
// 	type args struct {
// 		s *echotron.User
// 	}

// 	tests := []struct {
// 		name        string
// 		prep        func(f *fields)
// 		args        args
// 		expectedErr error
// 	}{
// 		{
// 			name: "nil sender",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(OnlyRealUsersText, gomock.Any(), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, nil)
// 			},
// 			args: args{
// 				s: nil,
// 			},
// 			expectedErr: domain.ErrMessageSenderNotSpecified,
// 		},
// 		{
// 			name: "nil sender, error notifying user",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(OnlyRealUsersText, gomock.Any(), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, domain.ErrInternal)
// 			},
// 			args: args{
// 				s: nil,
// 			},
// 			expectedErr: domain.ErrInternal,
// 		},
// 		{
// 			name: "empty username",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(OnlyUsersWithUsernameText, gomock.Any(), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, nil)
// 			},
// 			args: args{
// 				s: &echotron.User{},
// 			},
// 			expectedErr: domain.ErrEmptyUsername,
// 		},
// 		{
// 			name: "empty username, error notifying user",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(OnlyUsersWithUsernameText, gomock.Any(), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, domain.ErrInternal)
// 			},
// 			args: args{
// 				s: &echotron.User{},
// 			},
// 			expectedErr: domain.ErrInternal,
// 		},
// 		{
// 			name: "valid sender",
// 			prep: func(f *fields) {},
// 			args: args{
// 				s: &echotron.User{
// 					Username: "test",
// 				},
// 			},
// 			expectedErr: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			f := &fields{
// 				nil,
// 				mock_applicationbot.NewMockBotAPI(ctrl),
// 				mock_applicationbot.NewMockLogger(ctrl),
// 			}
// 			if tt.prep != nil {
// 				tt.prep(f)
// 			}

// 			b := &ApplicationBot{
// 				1,
// 				1,
// 				f.as,
// 				&domain.Application{},
// 				InitialState,
// 				f.l,
// 				&config.Config{},
// 				f.api,
// 			}
// 			err := b.setSenderInfo(tt.args.s)
// 			assert.Equal(t, tt.expectedErr, err)
// 		})
// 	}
// }

// func TestApplicationBot_parseUpdateText(t *testing.T) {
// 	type args struct {
// 		text string
// 	}
// 	tests := []struct {
// 		name         string
// 		prep         func(f *fields)
// 		args         args
// 		expectedText string
// 		expectedErr  error
// 	}{
// 		{
// 			name: "empty message",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(EnterUnemptyMessageText, gomock.Any(), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, nil)
// 			},
// 			args:         args{},
// 			expectedText: "",
// 			expectedErr:  domain.ErrEmptyMessageText,
// 		},
// 		{
// 			name: "message with only whitespaces",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(EnterUnemptyMessageText, gomock.Any(), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, nil)
// 			},
// 			args: args{
// 				text: "   \n",
// 			},
// 			expectedText: "",
// 			expectedErr:  domain.ErrEmptyMessageText,
// 		},
// 		{
// 			name: "empty message, err notifying",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(EnterUnemptyMessageText, gomock.Any(), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, domain.ErrInternal)
// 			},
// 			args:         args{},
// 			expectedText: "",
// 			expectedErr:  domain.ErrInternal,
// 		},
// 		{
// 			name: "valid message",
// 			prep: func(f *fields) {},
// 			args: args{
// 				text: "\t   /start   \n",
// 			},
// 			expectedText: "/start",
// 			expectedErr:  nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			f := &fields{
// 				nil,
// 				mock_applicationbot.NewMockBotAPI(ctrl),
// 				nil,
// 			}
// 			if tt.prep != nil {
// 				tt.prep(f)
// 			}

// 			b := &ApplicationBot{
// 				1,
// 				1,
// 				f.as,
// 				&domain.Application{},
// 				InitialState,
// 				f.l,
// 				&config.Config{},
// 				f.api,
// 			}
// 			got, err := b.parseUpdateText(tt.args.text)
// 			if tt.expectedErr != nil {
// 				assert.Equal(t, tt.expectedErr, err)
// 				return
// 			}
// 			assert.Equal(t, nil, err)
// 			assert.Equal(t, tt.expectedText, got)
// 		})
// 	}
// }

// func TestApplicationBot_processMessageUpdate(t *testing.T) {
// 	type args struct {
// 		u *echotron.Update
// 	}
// 	type results struct {
// 		t  string
// 		id int
// 	}
// 	tests := []struct {
// 		name        string
// 		prep        func(f *fields)
// 		args        args
// 		expected    results
// 		expectedErr error
// 	}{
// 		{
// 			name:        "nil update",
// 			prep:        func(f *fields) {},
// 			args:        args{},
// 			expected:    results{},
// 			expectedErr: domain.ErrInvalidUpdate,
// 		},
// 		{
// 			name: "nil message",
// 			prep: func(f *fields) {},
// 			args: args{
// 				u: &echotron.Update{},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrNoMessage,
// 		},
// 		{
// 			name: "invalid message chat",
// 			prep: func(f *fields) {},
// 			args: args{
// 				u: &echotron.Update{
// 					Message: &echotron.Message{},
// 				},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrChatNotPrivate,
// 		},
// 		{
// 			name: "invalid message sender",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(OnlyRealUsersText, gomock.Any(), gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				u: &echotron.Update{
// 					Message: &echotron.Message{
// 						Chat: echotron.Chat{
// 							Type: TgPrivateChatType,
// 						},
// 					},
// 				},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrMessageSenderNotSpecified,
// 		},
// 		{
// 			name: "invalid message text",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(EnterUnemptyMessageText, gomock.Any(), gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				u: &echotron.Update{
// 					Message: &echotron.Message{
// 						Chat: echotron.Chat{
// 							Type: TgPrivateChatType,
// 						},
// 						From: &echotron.User{
// 							Username: "test",
// 							ID:       1,
// 						},
// 					},
// 				},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrEmptyMessageText,
// 		},
// 		{
// 			name: "valid update",
// 			prep: func(f *fields) {},
// 			args: args{
// 				u: &echotron.Update{
// 					Message: &echotron.Message{
// 						ID: 1,
// 						Chat: echotron.Chat{
// 							Type: TgPrivateChatType,
// 						},
// 						From: &echotron.User{
// 							Username: "test",
// 							ID:       1,
// 						},
// 						Text: "   test   ",
// 					},
// 				},
// 			},
// 			expected: results{
// 				t:  "test",
// 				id: 1,
// 			},
// 			expectedErr: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			f := &fields{
// 				nil,
// 				mock_applicationbot.NewMockBotAPI(ctrl),
// 				nil,
// 			}
// 			if tt.prep != nil {
// 				tt.prep(f)
// 			}

// 			b := &ApplicationBot{
// 				1,
// 				1,
// 				f.as,
// 				&domain.Application{},
// 				InitialState,
// 				f.l,
// 				&config.Config{},
// 				f.api,
// 			}
// 			gotT, gotID, err := b.processMessageUpdate(tt.args.u)
// 			if tt.expectedErr != nil {
// 				assert.Equal(t, tt.expectedErr, err)
// 				return
// 			}
// 			assert.Equal(t, nil, err)
// 			assert.Equal(t, tt.expected.t, gotT)
// 			assert.Equal(t, tt.expected.id, gotID)
// 		})
// 	}
// }

// func TestApplicationBot_processCallbackQueryUpdate(t *testing.T) {
// 	type args struct {
// 		u *echotron.Update
// 	}
// 	type results struct {
// 		t  string
// 		id int
// 	}
// 	tests := []struct {
// 		name        string
// 		prep        func(f *fields)
// 		args        args
// 		expected    results
// 		expectedErr error
// 	}{
// 		{
// 			name:        "nil update",
// 			prep:        func(f *fields) {},
// 			args:        args{},
// 			expected:    results{},
// 			expectedErr: domain.ErrInvalidUpdate,
// 		},
// 		{
// 			name: "nil callback query",
// 			prep: func(f *fields) {},
// 			args: args{
// 				u: &echotron.Update{},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrInvalidCallbackQuery,
// 		},
// 		{
// 			name: "nil callback query message",
// 			prep: func(f *fields) {},
// 			args: args{
// 				u: &echotron.Update{
// 					CallbackQuery: &echotron.CallbackQuery{},
// 				},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrNoCallbackQueryMessage,
// 		},
// 		{
// 			name: "invalid message chat",
// 			prep: func(f *fields) {},
// 			args: args{
// 				u: &echotron.Update{
// 					CallbackQuery: &echotron.CallbackQuery{
// 						Message: &echotron.Message{},
// 					},
// 				},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrChatNotPrivate,
// 		},
// 		{
// 			name: "invalid message sender",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(OnlyRealUsersText, gomock.Any(), gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				u: &echotron.Update{
// 					CallbackQuery: &echotron.CallbackQuery{
// 						Message: &echotron.Message{
// 							Chat: echotron.Chat{
// 								Type: TgPrivateChatType,
// 							},
// 						},
// 					},
// 				},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrMessageSenderNotSpecified,
// 		},
// 		{
// 			name: "invalid message text",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(EnterUnemptyMessageText, gomock.Any(), gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				u: &echotron.Update{
// 					CallbackQuery: &echotron.CallbackQuery{
// 						From: &echotron.User{
// 							Username: "test",
// 							ID:       1,
// 						},
// 						Message: &echotron.Message{
// 							Chat: echotron.Chat{
// 								Type: TgPrivateChatType,
// 							},
// 						},
// 					},
// 				},
// 			},
// 			expected:    results{},
// 			expectedErr: domain.ErrEmptyMessageText,
// 		},
// 		{
// 			name: "valid update",
// 			prep: func(f *fields) {},
// 			args: args{
// 				u: &echotron.Update{
// 					CallbackQuery: &echotron.CallbackQuery{
// 						From: &echotron.User{
// 							Username: "test",
// 							ID:       1,
// 						},
// 						Message: &echotron.Message{
// 							ID: 1,
// 							Chat: echotron.Chat{
// 								Type: TgPrivateChatType,
// 							},
// 						},
// 						Data: "  test  ",
// 					},
// 				},
// 			},
// 			expected: results{
// 				t:  "test",
// 				id: 1,
// 			},
// 			expectedErr: nil,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			f := &fields{
// 				nil,
// 				mock_applicationbot.NewMockBotAPI(ctrl),
// 				nil,
// 			}
// 			if tt.prep != nil {
// 				tt.prep(f)
// 			}

// 			b := &ApplicationBot{
// 				1,
// 				1,
// 				f.as,
// 				&domain.Application{},
// 				InitialState,
// 				f.l,
// 				&config.Config{},
// 				f.api,
// 			}
// 			gotT, gotID, err := b.processCallbackQueryUpdate(tt.args.u)
// 			if tt.expectedErr != nil {
// 				assert.Equal(t, tt.expectedErr, err)
// 				return
// 			}
// 			assert.Equal(t, nil, err)
// 			assert.Equal(t, tt.expected.t, gotT)
// 			assert.Equal(t, tt.expected.id, gotID)
// 		})
// 	}
// }

// func TestApplicationBot_handleTextMessage(t *testing.T) {
// 	type args struct {
// 		t  string
// 		id int
// 	}
// 	tests := []struct {
// 		name          string
// 		prep          func(f *fields)
// 		args          args
// 		expectedState ConversationState
// 		expectedErr   error
// 	}{
// 		{
// 			name: "empty message",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(PressStartText, gomock.Any(), gomock.Any()).Times(1)
// 			},
// 			args:          args{},
// 			expectedState: InitialState,
// 			expectedErr:   nil,
// 		},
// 		{
// 			name: "approve command",
// 			prep: func(f *fields) {},
// 			args: args{
// 				t:  approveCommandPrefix,
// 				id: 1,
// 			},
// 			expectedState: InitialState,
// 			expectedErr:   domain.ErrUnauthorized,
// 		},
// 		{
// 			name: "reject command",
// 			prep: func(f *fields) {},
// 			args: args{
// 				t:  rejectCommandPrefix,
// 				id: 1,
// 			},
// 			expectedState: InitialState,
// 			expectedErr:   domain.ErrUnauthorized,
// 		},
// 		{
// 			name: "start command success",
// 			prep: func(f *fields) {
// 				f.as.EXPECT().GetApplicationByTgUserID(gomock.Any(), gomock.Any()).Times(1).Return(nil, domain.ErrApplicationNotFound)
// 				f.api.EXPECT().SendMessage(IntroductionText, gomock.Any(), gomock.Any()).Times(1)
// 				f.api.EXPECT().SendMessage(FullNamePromptText, gomock.Any(), gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				t:  startCommand,
// 				id: 1,
// 			},
// 			expectedState: WaitingForName,
// 			expectedErr:   nil,
// 		},
// 		{
// 			name: "start command fail",
// 			prep: func(f *fields) {
// 				f.as.EXPECT().GetApplicationByTgUserID(gomock.Any(), gomock.Any()).Times(1).Return(nil, nil)
// 				f.api.EXPECT().SendMessage(ApplicationAlreadyExistsText, gomock.Any(), gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				t:  startCommand,
// 				id: 1,
// 			},
// 			expectedState: InitialState,
// 			expectedErr:   domain.ErrApplicationAlreadyExists,
// 		},
// 		{
// 			name: "cancel command success",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(CancelText, gomock.Any(), gomock.Any()).Times(1)
// 			},
// 			args: args{
// 				t:  cancelCommand,
// 				id: 1,
// 			},
// 			expectedState: InitialState,
// 			expectedErr:   nil,
// 		},
// 		{
// 			name: "cancel command fail",
// 			prep: func(f *fields) {
// 				f.api.EXPECT().SendMessage(CancelText, gomock.Any(), gomock.Any()).Times(1).Return(echotron.APIResponseMessage{}, domain.ErrInternal)
// 			},
// 			args: args{
// 				t:  cancelCommand,
// 				id: 1,
// 			},
// 			expectedState: InitialState,
// 			expectedErr:   domain.ErrInternal,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			f := &fields{
// 				mock_applicationbot.NewMockApplicationService(ctrl),
// 				mock_applicationbot.NewMockBotAPI(ctrl),
// 				nil,
// 			}
// 			if tt.prep != nil {
// 				tt.prep(f)
// 			}

// 			b := &ApplicationBot{
// 				1,
// 				2,
// 				f.as,
// 				&domain.Application{},
// 				InitialState,
// 				f.l,
// 				&config.Config{},
// 				f.api,
// 			}
// 			err := b.handleTextMessage(tt.args.t, tt.args.id)
// 			assert.Equal(t, tt.expectedErr, err)
// 			assert.Equal(t, tt.expectedState, b.currentState)
// 		})
// 	}
// }
