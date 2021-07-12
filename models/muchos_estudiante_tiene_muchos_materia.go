package models

type MuchosEstudianteTieneMuchosMateria struct {
	IdEstudiante   *Estudiante `orm:"column(id_estudiante);rel(fk)"`
	IdMateria      *Materia    `orm:"column(id_materia);rel(fk)"`
	Nota1          float64     `orm:"column(nota_1);null"`
	Nota2          float64     `orm:"column(nota_2);null"`
	Nota3          float64     `orm:"column(nota_3);null"`
	NotaDefinitiva float64     `orm:"column(nota_definitiva);null"`
}
