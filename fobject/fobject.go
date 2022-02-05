package fobject

import (
	"fmt"
	"guid"
)

type FObject struct {
	this *FObject
	Name string
	guid guid.GUID

	onEveryTickDelegates []FDelegate
}

type FDelegate func()

func (this *FObject) Equals(other *FObject) bool {
	return this == other
}

func New(name string, onCreationCalls ...FDelegate) *FObject {
	temp, errorCode := guid.NewV4()

	if errorCode != nil {
		fmt.Printf("errorCode.Error(): %v\n", errorCode.Error())
		return nil
	}

	myFObject := &FObject{
		Name: name,
		guid: temp,
	}

	myFObject.this = myFObject

	myFObject.onInstanceCreated(&onCreationCalls, len(onCreationCalls))
	return myFObject
}

func (this *FObject) SubscribeOnEveryTick(onEveryTickCalls ...FDelegate) {
	if this.onEveryTickDelegates == nil {
		this.onEveryTickDelegates = onEveryTickCalls
	} else {
		this.onEveryTickDelegates = append(this.onEveryTickDelegates, onEveryTickCalls...)
	}
}
func (this *FObject) UnsubscribeOnEveryTick(onEveryTickCalls ...FDelegate) {
	if this.onEveryTickDelegates == nil {
		return
	} else {

	}
}

func (this *FObject) onInstanceCreated(delegates *[]FDelegate, size int) {
	for i := 0; i < size; i++ {
		(*delegates)[i]()
	}
}

func (this *FObject) onEveryTick() {
	for i := 0; i < len(this.onEveryTickDelegates); i++ {
		this.onEveryTickDelegates[i]()
	}
}

func (this *FObject) onInstanceDestroyed(delegates *[]FDelegate, size int) {
	for i := 0; i < size; i++ {
		(*delegates)[i]()
	}
}
