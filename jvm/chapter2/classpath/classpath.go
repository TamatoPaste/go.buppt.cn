package classpath

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	// TO-DO
	return nil
}
func (classpath *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	//TO-DO
	return nil, nil, nil
}
func (classpath *Classpath) String() string {
	//TO-DO
	return ""
}
