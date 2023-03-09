package generation

func GeneratePassword(int amountOfChar) string
{
	rand.Seed(time.Now().UnixNano())

    // define the length of the random string
    amountOfChar := amountOfChar

    // create a byte slice with length n
    b := make([]byte, amountOfChar)

    // fill the byte slice with random ASCII values
    for i := range b {
        b[i] = byte(rand.Intn(126-33) + 33)
    }

	return string(b)
}