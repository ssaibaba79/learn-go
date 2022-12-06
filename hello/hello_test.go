 package hello

import (
  "testing"
  )

type scenario struct {
    name string
    input [] string
    expected string
    result string
  }

  func (s scenario) test(t *testing.T) {
    s.result = SayHello(s.input)
    if s.expected != s.result {
      t.Errorf("\nScenario :'%s' FAILED. \n\tExpected: [%s]\n\tGot:[%s]\n",s.name, s.expected, s.result)
    }
  }

func TestSayHello(t *testing.T) {

  scenarios := []scenario{
    {
    name : "Empty input",
    input : [] string {},
    expected : "hello " ,
   },{
    name : "Single value input",
    input : [] string {"Anicha"},
    expected : "hello Anicha" ,
  },{
    name : "Multivalue input",
    input : [] string {"Anicha", "Sashti", "Poorni"},
    expected : "hello Anicha,Sashti,Poorni" ,
  }}

  for _, s := range scenarios {
    s.test(t)
  }
  
}