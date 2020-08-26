package main

import "testing"

func TestGreeting( t *testing.T ) {
	html := greeting("Test String")
	if(html != "<b>Test String</b>"){
		t.Errorf("Expected <b>Test String</b> but got %v",html)
	}
}
