package tool

var modelTemplate = `

func ({{obj}} *{{class}}) AfterCreate(tx *gorm.DB) (err error) {
	return
}

func ({{obj}} *{{class}}) BeforeCreate(tx *gorm.DB) (err error) {

	return
}
func ({{obj}} *{{class}}) BeforeUpdate(tx *gorm.DB) (err error) {
/*	if(this.ID > 0){
		changes := Changed(tx)
		if (len(changes) > 0) {
		  	this.LastModified =  time.Now()
		}
	}*/
	return
}

`
