package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
	"testing"
)

func Test_SendMessage(t *testing.T) {
	tests := []struct{
		name string
		tgBot interfaces.Notifications
		msg entity.Message
		wantErr bool
	}{
		{
			name: "send message over 4096 symbols",
			tgBot: NewTgBot("5471768780:AAEAbreeE6DDECknHmMrlD2Mfvedb5GIQ-w"),
			msg: entity.Message{
				Text: "hltyn5EGDstFdjHBvq12cUDidTDs23tcXGUEXXCSzVPuckb07ypuR22ZQj4IpGesNaHknz25jyGN9v02UNGzascFyMH8H83BSvbQ1vWhaYDKvewoK7jMJjTIAGkDz77t6s8LZAA6SvUIU3CFg0Qi0Xik3ouYblzoEr0u9f473SVxgjggPpvTjpUU5E0IsgZw78qBmLsGUmRdfHg4KKIE3LSsUlu14UhIjFzJxieLYGw6jsMoUTfTXOwS8Jtyxpi1SveyOo9iMAnRZgUkyJkIKQAHvDSw9q2KYx3znA5A7yeWfP4zm7OErd6bGXhdmHOzVcntx7jHJ0EsY1s2R5vhrwp5Pk1aF03N5mpZWDriuOLbPz84HWTVLpe3GrICi308q1RSN7hUwaf5sAoBQ3H6tJhlqRV6Gcie736NJLlxRAHDwSDnZ4z27mbTf2TIE8zjT0XoOyS0vQmfEASS289iaJCTejb9ZnNfn3mzoBDKEA27UvRIU8Hin711TGkDStN4NyKpBrWe3UvgKLGD2Z3IhSsY4rSYNH2ohcEUYV1bXuatwYcuVRa6IS5zyJd0iF76om6MYpeiWc24RtTkfYGOH1FJ0BvUH4z7P495nx8dzFv9QxSKtNmjxB0RpUBhfwjaheluCCU6A2xu82qjKk3aoNTsF9fwZx9srkz6kUhomcjgzFEeVeCAPDSWqD7FyJsWBgszPd39lgf03IMKzLDpKgqIyIbMHozMLh0bPKeeAplQj5P0FgJqOc03FjhIXB9UGtxIsRWVc7n1ZhK3o8DKc481gYljghUAXHaM3NLLD7v1L7icTUkWUa3cWYJeM0lHKItz1eayMbFGZ4lvHGeb33ofGefsj3hxMHRwPsafjzLmnflpws3QBD8Tw5OyPWa7UJlZmSIK0ZPMjXt9SufKN1R6Lr4CEmOvV432Z8320B6evVKI8cxblFfNcLZ8Gj1BJq0mt2jMnd2rZcxNlO4tFWRVrJygfTV51hooG3ZEguu6pE3DxlSLKy3d2cSHr1UbAaThpWyzyKcL0caTxP9npiMubiMuBboy8TUfFyCs8W78aJk2ICvjhQCaXzBGXGA9utBCFDYOVMrZhjO4cxSvTQ66yj6jpQyFHpDXdv2KWbq0UBZB8WN9e4YUoOIIsxcw8KoCZlkhVOuId0MyH5muKRE5vlotrR5eitquZhaVXtavJHproVlAGGpEcYx44XqI51eOd0ycj6edIkNXZ9sED9WGaYff2bcEAjsLiKT9o1aEEWB4t508WrpugtgFx4ytHjsYQsaMZZI8Zh2PIMvppiviGnuqrrNOdsYnUv1l3bNYACm2z99Cvg3URRqT7a4AbxQqby1Qwo3U45i2Zh71wvsVG6pQQCtuK6fdscMiVA461uBaUtc5eOmAQvVKsM8GVKtg0KUSdP88DEvY1Ri7GvoTkGCoxxXRrabOc6fw6ZUSiZusDnBkRnqPZ2vCI0jfTS5aVKRvX9qA5TisCvTPEgKiA2yVP2Mb5xzpGbtNZHkxaBEOULzRbCPzjVVbiHIrzoT2vJkjTU5fHgfOmrqpkwadTpjTWLR8loFXjKSzQG3L8sXN0tK2AWnbzdvuNuawyYXZmAAH2p22wEODpM5j1kBmfRhCJTbwTvrQFs99i3Yr04sgvCBLJC0YwaoFqNdbbyFWB59jsqhYmuP7cgQF3uKL7g89cPwI1ZsRXlqNi1fk1qiWjRUHswLEwmjJPY7zLdeM4OYwvq8hiNmol9ESJT0JMcmwD3jy2B8aN6DMXN185xNFeNaoOlRKsA4rjzEYqjUsDLb0O8SyR5vs8SICdOhm6NgijICNr8v4fzgMVLGlWOPyMahcIZsUgcpjPodEfdOeF3lqLESbz2FThqoGzn6OyaPjHXodVjSh6gKCyCf5VPDvUPLUvogzC3mYsIGnJyEkC5WZ2ZnE6t2S4BZB9VkLktYIytK8b3eCQvmMNCArpQ4rNMf7vliQKZPEElE6Xg2jHJACRmo3wOBmvlQq6Q5v7imQeeVe66JniaiuMwpaC440m1tsHE8U01KDTh8VLjObwRqh4MEuRc3mnPt7e874oZUx5mjfAmlns8UcC6ZnSWmoQ6xVZpwfKK3eNpASvJLT3gbl5yLgXsbWJkq7fsrXu7N9IeFc61xxJd2oi4MpnPI2qPdFsVkVFG3PFPzYiN1iq3imFS533L7SQd2gWIrg7Rryl9RSKn8ELRXIUye5JTL5yPRmAeoMbPxWQ6zmr5lRkxLjOQYRdEo8uxP3HBQ1Be0uQ7ose6phaxCVEeTj42ZzKFTiHjt67efxh7wCeZB5TqQlXRm8gcEaQ8fggWZaTkc48vNllcwGWgv4qjGnfNwzgu1gZERhJjRUhD6k9vG9kyACCcYrA1X6lR4erDqfDb4qXIUp0yLME6fwsWX9aN8IpbjJdKS3PeeJJBJ6zPsgyiKVXcfIXmyn0XWxOeH74ZK6cUoEczpErVwuB7qkWIpaxhajElYjyphMNhU5qh976iCphoRmTK8EBjADYU0yLBd4eiQIrZft9WCnT6OJ1UjhoZeA5nNQsnVI5gdBG2HpNXmLg9Jielb6jupPxfqHcAP1TAt6L8oGe852bRQLbU91xDAm41Mo7GN1pGsQn08wFfsoxqtUgGzDWX6QwMAoLfAOJrmCoaEEaZKn2fJeZVd8sGeyv3qPIhwI6KIi5Yv4W5eiTndSDlQDr1mju3IPMNQyGthfPlhv4P4cgruKXQMThZlD7M3YbwDeJwc7dBAyrnJmXtNFq5IwoHflY67u6PdM5B8oh9gWZz4eQ7p0rc3LPZvxhMgRl1wmYHCu1y33BqixKX54TXzs9Pp5YbMXjSy2vFzhBPl7ie7a5IehVaIwYJH5ncB7AfICzhvdlBvIAZ3gp7EJPFNGPWjUZJ4NFA0cwKFDnROLFzTAYAOkfEFuUd8KlngcqYiIxcNozRr9hHpv7rXJOhB7B4cYI5lngcJPVDxUXLNALNuDFSTUYg6zGRBxi8x6bW62dR8FxCXRPh8eT0S7lLlFfreZQFeCN86TOwrrJDVsFhocKkTJcrm22TR7Ackb6pwmU71HHlW34OngbZLop8WWU12SpwOOZMZl9GMZSSEXqiNOFZA2ROW2nA1zD6Gz2cSS8Z70xFI5xFtc0a3pSTQIczWWi4SxSMV00JQq8cfykOjQhYWgtk7tUibsnWzmpRb3rzQ0icB4jj5TqhYFScF0RE1ieUqfSRITcTznBsi5Rv0sYub2ahYuuhcCUpGQErYsDzflfw9MgJx9jCwfSRYKPeSNDxo8kJ0OA4S4zGY7vCrKldyXiNykHpoQPmg3lIWMFupMKxNALQ1TjwHR4qXCmsfG5LN69KQoHxY2p2uAc96CTUZ7MwvN2T2vR92Ge7K3M8sNqqMuQ6x3ZYFkSRwIOZlygKcGqESNdlx7uQ1t0vSNYm8T5B2eeCfwJZDbUpvxpa2CcoMZW3cTfGSe9Fd2jhcdEXsAoGvFBO2jzJxeOepwnWp8bhzrKQY0FaDF7Nci9bpw0hriOgCKnI3ADr1Fk6QhbSNavuIQDWnmg4A2g55ntW8DHWLUDqUQYytbLhu5flIXsmC3VbIiOBg3MCZrctC987hOgpXaE178jHCgtTpZMpYt5S7WmJutX5tf6TWa3IbEPaP11eInBIOFmgu37h14q2kxTEVxmVUEuBbf3GazPXeGDb1cWIZMp2Z9AqjWBtFjfpG18rktWDX5Yk5p3eEqe5fAdPwbfeCpFFBOHFwBqazkxVDzOy5PwJrXRBdfTWCnHjjkQ4GyLU33Uuk2Z7RMilzUDkEpRsp8plKH9pPNtqXhLLp23FCx0nIpX745Lmh5e8npfafvwX3jXUT4lyKKC6BZARmuTTZldng6y1GmUigeKc899vTeaxTz0uqwhFeyg3tBOrmCF2D1Bjk0DRgtnzO0VYJvpRmghUK0yBlEzEu526nqVKSTHkOlWXvUQZzKkBUy3pHGuFiEl2umW3BMJlqliMt0K1J7qjjRJTQa50qaPSC5sQi7pWyeEZxZhEarr0he52xugcUH4SccD0Ehym18EYwZFbxoXp0YYqPpOxwwZnEz2Q1CmCKQRtGjLiCSGPEu",
				To: "1028417962",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tgBot.SendMessage(tt.msg); (err != nil) != tt.wantErr {
				t.Errorf("sendmessage() error = %v, logs %v", err, tt.wantErr)
			}
		})
	}
}