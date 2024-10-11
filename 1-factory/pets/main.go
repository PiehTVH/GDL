package pets

type IPet interface {
	GetName() string
	GetSound() string
	GetAge() int
}

func PetFactory(petType string) IPet {
	if petType == "dog" {
		return &dog{
			pet{
				name:  "Chester",
				age:   2,
				sound: "bark",
			},
		}
	} else if petType == "cat" {
		return &cat{
			pet{
				name:  "Mr. Buttons",
				age:   3,
				sound: "meow",
			},
		}
	}

	return nil
}
