package services

import (
	"context"
	"server_simple/internal/models"
	"time"
)

type GAME_READY_START_DATA struct {
	Event int `json:"event"`
	Data  struct {
		CountdownSecond int    `json:"countdown_second"`
		Question        string `json:"question"`
		QuestionTips    string `json:"question_tips"`
		Status          int    `json:"status"`
	} `json:"data"`
}

type EVENT_GAME_ANSWER_DATA struct {
	Event int `json:"event"`
	Data  struct {
		Avatar       string `json:"avatar"`
		Nickname     string `json:"nickname"`
		Answer       string `json:"answer"`
		AnswerResult int    `json:"answer_result"` //1.作答正确，0不正确
		Sec          int    `json:"sec"`
		Uid          int    `json:"uid"`
	} `json:"data"`
}

func (this *gdrawWs) countdowner(ctx context.Context, countdownSec int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case c <- countdownSec:
				countdownSec--
			}
			time.Sleep(time.Second * 1)
		}
	}()
	return c
}

func (this *gdrawWs) questionEmiter(roomId string) (question models.GdrawQuestion) {
	roomQuestion, _ := models.GdrawRoomQuestioner.GetWithRoomStatus(roomId, 0)
	question.Id = roomQuestion.QuestionId
	question.Question = roomQuestion.Question
	question.QuestionTips = roomQuestion.QuestionTips
	_ = models.GdrawRoomQuestioner.UpdateWithId(roomQuestion.Id)
	return
}
