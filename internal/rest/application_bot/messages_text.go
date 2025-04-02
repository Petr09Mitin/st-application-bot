package applicationbot

import (
	"fmt"

	"github.com/st-application-bot/domain"
)

const (
	SubscriptionBotLink = `https://t.me/SevTribunaPodpiska_bot`

	IntroductionText = `Дорогой Брат! 

Этот бот создан для подачи заявки в Московский состав «Северной Трибуны». Просим тебя ответить на вопросы ниже. 

Прежде чем подавать заявку, пожалуйста, ознакомься с правилами: https://telegra.ph/Pravila-povedeniya-na-meropriyatiyah-Severnoj-Tribuny-05-27

Обрати внимание, что участие в активе подразумевает ежемесячный взнос (от 350р) на нужды наших братьев, находящихся на фронте! Взнос вносится в формате ежемесячной подписки на чат актива. После одобрения заявки администратором тебе будет предоставлена ссылка на бота для оформления подписки. 

<i>Чтобы очистить свою заявку, нажми /cancel.</i>`

	SecondRosterIntroducationText = `Дорогой Брат! 

Этот бот создан для подачи заявки в Московский состав «Северной Трибуны». Просим тебя ответить на вопросы ниже. 

Прежде чем подавать заявку, пожалуйста, ознакомься с правилами: https://telegra.ph/Pravila-povedeniya-na-meropriyatiyah-Severnoj-Tribuny-05-27

Обрати внимание, что нахождение в составе подразумевает регулярное посещение турниров и репетиций "Северной Трибуны"!

В чате будет доступна актуальная информация о посещаемых мероприятиях и билеты.

P. S. Если по объективным причинам ты не можешь делиться своими личными данными (например, служба в силовых структурах и т. д.) свяжись с @stadmin001

<i>Чтобы очистить свою заявку, нажми /cancel.</i>`

	FullNamePromptText = `<b>Пожалуйста, представься - пришли свои ФИО полностью.</b>

<i>Пример: Иванов Иван Иванович</i>`
	PhonePromptText = `<b>Укажи свой контактный номер телефона</b>

<i>Пример: +71112223344</i>`
	DistrictPromptText = `<b>Укажи свой район (либо город) проживания</b> 

<i>Примеры: Выхино, Химки</i>`
	PurposePromptText     = `<b>С какой целью ты пришел в актив «Северной Трибуны»?</b>`
	SocialMediaPromptText = `<b>Пожалуйста, отправь ссылку на свою страницу в какой-либо социальной сети (ВКонтакте, Одноклассники или иной).

Если тебя нет ни в одной социальной сети, введи "-".</b>`
	IsSevChelParticipantPromptText = `<b>Являешься ли ты участником объединения "Северный Человек"?</b>`
	InformationSourcePromptText    = `<b>Как ты узнал о Северной Трибуне?</b>`
	ConclusionTextFirstRoster      = `Спасибо за обратную связь!

Твоя заявка находится на рассмотрении. После одобрения заявки администратором мы пришлем ссылку на бот для оформления ежемесячной подписки/сбора.`
	ConclusionTextSecondRoster = `Спасибо за обратную связь!
	
Твоя заявка находится на рассмотрении. После одобрения заявки администратором ты будешь автоматически добавлен в чат актива.

Чтобы завершить процесс подачи заявки, перейди по ссылке ` + SecondRosterChatLink
	PressStartText           = `Чтобы начать, нажми /start.`
	CancelText               = `Заявка отменена. Чтобы попробовать еще раз, нажми /start.`
	NoMessageText            = `Пожалуйста, пришли текстовое сообщение либо нажми кнопку!`
	OnlyRealUsersText        = `Мы взаимодействуем только с реальными пользователями!`
	EnterUnemptyMessageText  = `Пожалуйста, введи непустое сообщение.`
	IncorrectPhoneFormatText = `<b>Неверный формат номера телефона.</b> 
	
<i>Пожалуйста, укажи номер в формате +71112223344.</i>`
	IncorrectFullNameFormatText = `<b>Неверный формат ФИО.</b>

<i>Укажи ФИО в формате: Иванов Иван Иванович</i>`
	InternalErrorText            = `Что-то пошло не так... Чтобы начать заново, нажми /start`
	ApplicationAlreadyExistsText = `Ты уже подавал заявку!`
	OnlyUsersWithUsernameText    = `Чтобы взаимодействовать с ботом, пожалуйста, установи уникальное имя пользователя!
<i>Уникальное имя этого бота - @north_tribune_applications_bot</i>`
	ApproveText = `Дорогой Брат! 

Твоя заявка одобрена, добро пожаловать в «Северную Трибуну».

Перейди в бот для оформления ежемесячной подписки ` + SubscriptionBotLink + ` и следуй всем инструкциям.

До встречи на трибуне✊🏻`
	SecondRosterChatLink    = `https://t.me/+TNNsR-dmWTcxYTQy`
	SecondRosterApproveText = `Дорогой Брат! 

Твоя заявка одобрена, добро пожаловать в «Северную Трибуну».

Если ты не был автоматически добавлен в чат, напиши администратору @stadmin001.

До встречи на трибуне✊🏻`
	RejectText = `К сожалению, твоя заявка отклонена администратором.`
)

func getIntroductionText(isSecondRoster bool) string {
	if isSecondRoster {
		return SecondRosterIntroducationText
	}

	return IntroductionText
}

func getApplicationMessageText(application *domain.Application) (text string) {
	var isSevChelMemberString string
	if application.IsSevChelParticipant {
		isSevChelMemberString = "Да"
	} else {
		isSevChelMemberString = "Нет"
	}

	return fmt.Sprintf(
		`
ID: %d
Ссылка на аккаунт: https://t.me/%s
ФИО: %s
Город/Район: %s
Ссылка на соцсеть: %s
Состоит в СЧ: %s
		`,
		application.ID,
		application.TelegramUsername,
		application.FullName,
		application.District,
		application.SocialMediaLink,
		isSevChelMemberString,
	)
}

func getConclusionText(isSecondRoster bool) string {
	if isSecondRoster {
		return ConclusionTextSecondRoster
	}

	return ConclusionTextFirstRoster
}

func getApproveText(isSecondRoster bool) string {
	if isSecondRoster {
		return SecondRosterApproveText
	}

	return ApproveText
}

func getApproveAdminMessageText(id uint64, tgUsername string) string {
	return fmt.Sprintf("Заявка с ID %d (@%s) одобрена!", id, tgUsername)
}

func getRejectAdminMessageText(id uint64, tgUsername string) string {
	return fmt.Sprintf("Заявка с ID %d (@%s) отклонена!", id, tgUsername)
}
