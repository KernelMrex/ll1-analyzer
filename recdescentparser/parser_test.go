package recdescentparser

import (
	"strings"
	"testing"
)

func TestSingleLine(t *testing.T) {
	s := `prog id var id : int begin read(id, id); id := id + ((--num)); end`
	r := strings.NewReader(s)
	if !Process(r) {
		t.Fail()
	}
}

func TestSingleLineCaseInsensitive(t *testing.T) {
	s := `PROG id vAr id : int bEgIn reAD(Id, Id); ID := id + ((--nUm)); eNd`
	r := strings.NewReader(s)
	if !Process(r) {
		t.Fail()
	}
}

func TestMultiLineCaseInsensitive(t *testing.T) {
	s := `
PROG    id
vAr
  id: int
bEgIn
  reAD(  Id,   Id  );reAD(Id);

  write(  Id,   Id  );
  write(Id);

  ID :=  id    +   (  (  -  -  nUm  )  );
  ID :=  id    *   (  (  -  -  nUm  )  );

  write(  Id,   Id  );

eNd
`
	r := strings.NewReader(s)
	if !Process(r) {
		t.Fail()
	}
}

func TestMultiplicationOperator(t *testing.T) {
	s := `
prog id
var
  id: int
begin
	id := id * id;
	id := id * (-id);
	id := id*(-id);
    id := id*id;
end
`
	r := strings.NewReader(s)
	if !Process(r) {
		t.Fail()
	}
}
