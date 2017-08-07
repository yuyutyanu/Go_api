package pack

func ErrHandler(err error)  {
	if err != nil{
		panic(err)
	}
}
