package state

import (
	"sync"
	"testing"
	"time"

	"github.com/mk6i/retro-aim-server/oscar"
	"github.com/stretchr/testify/assert"
)

func TestSession_SetAndGetAwayMessage(t *testing.T) {
	s := NewSession()
	assert.Empty(t, s.AwayMessage())

	msg := "here's my message"
	s.SetAwayMessage(msg)
	assert.Equal(t, msg, s.AwayMessage())
}

func TestSession_SetAndGetID(t *testing.T) {
	s := NewSession()
	// make sure NewSession creates a default ID
	assert.NotEmpty(t, s.SetID)
	newID := "new-id"
	s.SetID(newID)
	assert.Equal(t, newID, s.ID())
}

func TestSession_IncrementAndGetWarning(t *testing.T) {
	s := NewSession()
	assert.Zero(t, s.Warning())
	s.IncrementWarning(1)
	s.IncrementWarning(2)
	assert.Equal(t, uint16(3), s.Warning())
}

func TestSession_SetAndGetInvisible(t *testing.T) {
	s := NewSession()
	assert.False(t, s.Invisible())
	s.SetInvisible(true)
	assert.True(t, s.Invisible())
	s.SetInvisible(false)
	assert.False(t, s.Invisible())
}

func TestSession_SetAndGetScreenName(t *testing.T) {
	s := NewSession()
	assert.Empty(t, s.ScreenName())
	sn := "user-screen-name"
	s.SetScreenName(sn)
	assert.Equal(t, sn, s.ScreenName())
}

func TestSession_SendMessage(t *testing.T) {
	type fields struct {
		awayMessage string
		closed      bool
		id          string
		idle        bool
		idleTime    time.Time
		invisible   bool
		msgCh       chan oscar.SNACMessage
		mutex       sync.RWMutex
		screenName  string
		signonTime  time.Time
		stopCh      chan struct{}
		warning     uint16
	}
	type args struct {
		msg oscar.SNACMessage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SessSendStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				awayMessage: tt.fields.awayMessage,
				closed:      tt.fields.closed,
				id:          tt.fields.id,
				idle:        tt.fields.idle,
				idleTime:    tt.fields.idleTime,
				invisible:   tt.fields.invisible,
				msgCh:       tt.fields.msgCh,
				mutex:       tt.fields.mutex,
				screenName:  tt.fields.screenName,
				signonTime:  tt.fields.signonTime,
				stopCh:      tt.fields.stopCh,
				warning:     tt.fields.warning,
			}
			assert.Equalf(t, tt.want, s.RelayMessage(tt.args.msg), "RelayMessage(%v)", tt.args.msg)
		})
	}
}

func TestSession_SetAwayMessage(t *testing.T) {
	type fields struct {
		awayMessage string
		closed      bool
		id          string
		idle        bool
		idleTime    time.Time
		invisible   bool
		msgCh       chan oscar.SNACMessage
		mutex       sync.RWMutex
		screenName  string
		signonTime  time.Time
		stopCh      chan struct{}
		warning     uint16
	}
	type args struct {
		awayMessage string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				awayMessage: tt.fields.awayMessage,
				closed:      tt.fields.closed,
				id:          tt.fields.id,
				idle:        tt.fields.idle,
				idleTime:    tt.fields.idleTime,
				invisible:   tt.fields.invisible,
				msgCh:       tt.fields.msgCh,
				mutex:       tt.fields.mutex,
				screenName:  tt.fields.screenName,
				signonTime:  tt.fields.signonTime,
				stopCh:      tt.fields.stopCh,
				warning:     tt.fields.warning,
			}
			s.SetAwayMessage(tt.args.awayMessage)
		})
	}
}

func TestSession_TLVUserInfo(t *testing.T) {
	tests := []struct {
		name           string
		givenSessionFn func() *Session
		want           oscar.TLVUserInfo
	}{
		{
			name: "user is active and visible",
			givenSessionFn: func() *Session {
				s := NewSession()
				s.SetSignonTime(time.Unix(1, 0))
				s.SetScreenName("xXAIMUSERXx")
				s.IncrementWarning(10)
				return s
			},
			want: oscar.TLVUserInfo{
				ScreenName:   "xXAIMUSERXx",
				WarningLevel: 10,
				TLVBlock: oscar.TLVBlock{
					TLVList: oscar.TLVList{
						oscar.NewTLV(0x03, uint32(1)),
						oscar.NewTLV(0x01, uint16(0x0010)),
						oscar.NewTLV(0x06, uint16(0x0000)),
						oscar.NewTLV(0x04, uint16(0)),
						oscar.NewTLV(0x0D, capChat),
					},
				},
			},
		},
		{
			name: "user has away message set",
			givenSessionFn: func() *Session {
				s := NewSession()
				s.SetSignonTime(time.Unix(1, 0))
				s.SetAwayMessage("here's my away essage")
				return s
			},
			want: oscar.TLVUserInfo{
				TLVBlock: oscar.TLVBlock{
					TLVList: oscar.TLVList{
						oscar.NewTLV(0x03, uint32(1)),
						oscar.NewTLV(0x01, uint16(0x30)),
						oscar.NewTLV(0x06, uint16(0x0000)),
						oscar.NewTLV(0x04, uint16(0)),
						oscar.NewTLV(0x0D, capChat),
					},
				},
			},
		},
		{
			name: "user is invisible",
			givenSessionFn: func() *Session {
				s := NewSession()
				s.SetSignonTime(time.Unix(1, 0))
				s.SetInvisible(true)
				return s
			},
			want: oscar.TLVUserInfo{
				TLVBlock: oscar.TLVBlock{
					TLVList: oscar.TLVList{
						oscar.NewTLV(0x03, uint32(1)),
						oscar.NewTLV(0x01, uint16(0x0010)),
						oscar.NewTLV(0x06, uint16(0x0100)),
						oscar.NewTLV(0x04, uint16(0)),
						oscar.NewTLV(0x0D, capChat),
					},
				},
			},
		},
		{
			name: "user is idle",
			givenSessionFn: func() *Session {
				s := NewSession()
				s.SetSignonTime(time.Unix(1, 0))
				// now() returns T=1000 when SetIdle() is called
				s.nowFn = func() time.Time { return time.Unix(1000, 0) }
				s.SetIdle(1 * time.Second)
				// now() returns T=2000 when TLVUserInfo() is called
				s.nowFn = func() time.Time { return time.Unix(2000, 0) }
				return s
			},
			want: oscar.TLVUserInfo{
				TLVBlock: oscar.TLVBlock{
					TLVList: oscar.TLVList{
						oscar.NewTLV(0x03, uint32(1)),
						oscar.NewTLV(0x01, uint16(0x0010)),
						oscar.NewTLV(0x06, uint16(0x0000)),
						oscar.NewTLV(0x04, uint16(1001)),
						oscar.NewTLV(0x0D, capChat),
					},
				},
			},
		},
		{
			name: "user goes idle then returns",
			givenSessionFn: func() *Session {
				s := NewSession()
				s.SetSignonTime(time.Unix(1, 0))
				s.SetIdle(1 * time.Second)
				s.UnsetIdle()
				return s
			},
			want: oscar.TLVUserInfo{
				TLVBlock: oscar.TLVBlock{
					TLVList: oscar.TLVList{
						oscar.NewTLV(0x03, uint32(1)),
						oscar.NewTLV(0x01, uint16(0x0010)),
						oscar.NewTLV(0x06, uint16(0x0000)),
						oscar.NewTLV(0x04, uint16(0)),
						oscar.NewTLV(0x0D, capChat),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.givenSessionFn()
			assert.Equal(t, tt.want, s.TLVUserInfo())
		})
	}
}

func TestSession_SendAndRecvMessage_ExpectSessSendOK(t *testing.T) {
	s := NewSession()

	msg := oscar.SNACMessage{
		Frame: oscar.SNACFrame{
			FoodGroup: oscar.ICBM,
		},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer s.Close()
		status := s.RelayMessage(msg)
		assert.Equal(t, SessSendOK, status)
	}()

loop:
	for {
		select {
		case m := <-s.ReceiveMessage():
			assert.Equal(t, msg, m)
		case <-s.Closed():
			break loop
		}
	}

	wg.Wait()
}

func TestSession_SendMessage_SessSendClosed(t *testing.T) {
	s := Session{
		msgCh:  make(chan oscar.SNACMessage, 1),
		stopCh: make(chan struct{}),
	}
	s.Close()
	if res := s.RelayMessage(oscar.SNACMessage{}); res != SessSendClosed {
		t.Fatalf("expected SessSendClosed, got %+v", res)
	}
}

func TestSession_SendMessage_SessQueueFull(t *testing.T) {
	bufSize := 10
	s := Session{
		msgCh:  make(chan oscar.SNACMessage, bufSize),
		stopCh: make(chan struct{}),
	}
	for i := 0; i < bufSize; i++ {
		assert.Equal(t, SessSendOK, s.RelayMessage(oscar.SNACMessage{}))
	}
	assert.Equal(t, SessQueueFull, s.RelayMessage(oscar.SNACMessage{}))
}

func TestSession_Close_Twice(t *testing.T) {
	s := Session{
		stopCh: make(chan struct{}),
	}
	s.Close()
	s.Close() // make sure close is idempotent
	if !s.closed {
		t.Fatal("expected session to be closed")
	}
	select {
	case <-s.Closed():
	case <-time.After(1 * time.Second):
		t.Fatalf("channel is not closed")
	}
}

func TestSession_Close(t *testing.T) {
	s := NewSession()
	select {
	case <-s.Closed():
		assert.Fail(t, "channel is closed")
	default:
		// channel is open by default
	}
	s.Close()
	<-s.Closed()
}
