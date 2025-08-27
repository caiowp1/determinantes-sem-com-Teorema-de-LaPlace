package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprimeMatriz(mat [][]int) {
	var contI int
	var contJ int
	for contI = 0; contI < len(mat); contI++ {
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			fmt.Print(mat[contI][contJ], " ")
		}
		fmt.Println()
	}
}

func iniciaMatrizAleatoria(mat [][]int) {
	var contI int
	var contJ int
	var ordem int
	var limiteSuperior int
	var numAleatorio int

	ordem = len(mat)
	limiteSuperior = (ordem * ordem) / 2

	for contI = 0; contI < ordem; contI++ {
		for contJ = 0; contJ < ordem; contJ++ {
			numAleatorio = rand.Intn(limiteSuperior + 1)
			mat[contI][contJ] = numAleatorio
		}
	}
}

// Base do cálculo de determinantes

func verificaQuadradaOrdem(mat [][]int) (bool, int) {
	var numLinhas int
	var numColunas int
	var ehQuadrada bool

	numLinhas = len(mat)
	numColunas = len(mat[0])

	ehQuadrada = false
	if numLinhas == numColunas {
		ehQuadrada = true
	}

	return ehQuadrada, numLinhas
}

func calculaSinal(indiceL int, indiceC int) int {
	var sinal int

	sinal = -1
	if ((indiceL + indiceC) % 2) == 0 {
		sinal = 1
	}

	return sinal
}

func copiaMatrizMaiorParaMenor(maior [][]int, menor [][]int, isqn int, jsqn int) {
	var contAi, contAj, contBi, contBj, temp, numL, numC int
	numL = len(menor)
	numC = len(menor[0])

	contAi = 0
	for contBi = 0; contBi < numL; contBi++ {
		if contAi == isqn {
			contAi++
		}
		contAj = 0
		for contBj = 0; contBj < numC; contBj++ {
			if contAj == jsqn {
				contAj++
			}
			temp = maior[contAi][contAj]
			menor[contBi][contBj] = temp
			contAj++
		}
		contAi++
	}
}

func detOrdem1(mat [][]int) int {
	return mat[0][0]
}

func detOrdem2(mat [][]int) int {
	var diagonalP int
	var diagonalI int

	diagonalP = mat[0][0] * mat[1][1]
	diagonalI = mat[1][0] * mat[0][1]
	return (diagonalP - diagonalI)
}

func detOrdemNBase(mat [][]int) int {
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC, cont int
	var matMenor [][]int

	numL = len(mat)
	numC = len(mat[0])

	resposta = 0
	contL = 0

	for contC = 0; contC < numC; contC++ {
		cofator = mat[contL][contC]
		sinal = calculaSinal(contL, contC)
		matMenor = make([][]int, numL-1)

		for cont = 0; cont < (numL - 1); cont++ {
			matMenor[cont] = make([]int, numC-1)
		}

		copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
		detTemp, _ = determinanteBase(matMenor)
		resposta = resposta + (cofator * sinal * detTemp)
	}

	return resposta

}

func determinanteBase(mat [][]int) (int, int64) {
	var ordem int
	var ehQuadrada bool
	var det int
	var tempoInicio time.Time
	var tempoFim time.Time
	var tempoExecucao int64

	tempoInicio = time.Now()

	ehQuadrada, ordem = verificaQuadradaOrdem(mat)

	if ehQuadrada {
		switch ordem {
		case 1:
			det = detOrdem1(mat)
		case 2:
			det = detOrdem2(mat)
		default:
			det = detOrdemNBase(mat)
		}
	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}

	tempoFim = time.Now()

	tempoExecucao = tempoFim.UnixNano() - tempoInicio.UnixNano()

	return det, tempoExecucao
}

// Otimização do cálculo de determinantes

func escolheLinhaColunaComMaisZeros(mat [][]int) (bool, int) {
	var numL, numC int
	var usarLinha bool
	var indice int
	var zerosPorLinha []int
	var zerosPorColuna []int
	var contI, contJ, maxZeros int

	numL = len(mat)
	numC = len(mat)

	zerosPorLinha = make([]int, numL)
	zerosPorColuna = make([]int, numC)

	usarLinha = true
	indice = 0
	maxZeros = 0

	for contI = 0; contI < numL; contI++ {
		for contJ = 0; contJ < numC; contJ++ {
			if mat[contI][contJ] == 0 {
				zerosPorLinha[contI]++
				zerosPorColuna[contJ]++
			}
		}
	}

	for contI = 0; contI < numL; contI++ {
		if zerosPorLinha[contI] > maxZeros {
			maxZeros = zerosPorLinha[contI]
			indice = contI
			usarLinha = true
		}
	}
	for contJ = 0; contJ < numC; contJ++ {
		if zerosPorColuna[contJ] > maxZeros {
			maxZeros = zerosPorColuna[contJ]
			indice = contJ
			usarLinha = false
		}
	}

	return usarLinha, indice
}

func detOrdemNOtimizado(mat [][]int) int {
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC, cont int
	var matMenor [][]int
	var usarLinha bool
	var indiceFixo int

	numL = len(mat)
	numC = len(mat[0])
	resposta = 0

	usarLinha, indiceFixo = escolheLinhaColunaComMaisZeros(mat)

	if usarLinha {
		for contC = 0; contC < numC; contC++ {
			cofator = mat[indiceFixo][contC]
			if cofator != 0 {
				sinal = calculaSinal(indiceFixo, contC)

				matMenor = make([][]int, numL-1)
				for cont = 0; cont < (numL - 1); cont++ {
					matMenor[cont] = make([]int, numC-1)
				}

				copiaMatrizMaiorParaMenor(mat, matMenor, indiceFixo, contC)
				detTemp, _ = determinanteOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)
			}
		}
	} else {
		for contL = 0; contL < numL; contL++ {
			cofator = mat[contL][indiceFixo]
			if cofator != 0 {
				sinal = calculaSinal(contL, indiceFixo)

				matMenor = make([][]int, numL-1)
				for cont = 0; cont < (numL - 1); cont++ {
					matMenor[cont] = make([]int, numC-1)
				}

				copiaMatrizMaiorParaMenor(mat, matMenor, contL, indiceFixo)
				detTemp, _ = determinanteOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)
			}
		}
	}

	return resposta
}

func determinanteOtimizado(mat [][]int) (int, int64) {
	var ordem int
	var ehQuadrada bool
	var det int
	var tempoInicio time.Time
	var tempoFim time.Time
	var tempoExecucao int64

	tempoInicio = time.Now()

	ehQuadrada, ordem = verificaQuadradaOrdem(mat)

	if ehQuadrada {
		switch ordem {
		case 1:
			det = detOrdem1(mat)
		case 2:
			det = detOrdem2(mat)
		default:
			det = detOrdemNOtimizado(mat)
		}
	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}

	tempoFim = time.Now()

	tempoExecucao = tempoFim.UnixNano() - tempoInicio.UnixNano()

	return det, tempoExecucao
}

func main() {

	var determinante int
	var tempoExecucaoBase [3]int64
	var tempoExecucaoOtimizado [3]int64
	var somaBase, somaOtimizado int64
	var mediaBase, mediaOtimizado int64
	var numLinhas int
	var numColunas int
	var cont int
	var contOrdem int

	for contOrdem = 3; contOrdem <= 11; contOrdem += 2 {

		var matrix [][]int

		numLinhas = contOrdem
		numColunas = contOrdem

		somaBase = 0
		somaOtimizado = 0

		for cont = 0; cont < 3; cont++ {
			tempoExecucaoBase[cont] = 0
			tempoExecucaoOtimizado[cont] = 0
		}

		matrix = make([][]int, numLinhas)

		for cont = 0; cont < numLinhas; cont++ {
			matrix[cont] = make([]int, numColunas)
		}

		for cont = 0; cont < 3; cont++ {
			iniciaMatrizAleatoria(matrix)
			fmt.Println("Matriz Gerada Aleatoriamente de Ordem", contOrdem, "x", contOrdem)
			imprimeMatriz(matrix)

			_, tempoExecucaoBase[cont] = determinanteBase(matrix)
			determinante, tempoExecucaoOtimizado[cont] = determinanteOtimizado(matrix)
			somaBase += tempoExecucaoBase[cont]
			somaOtimizado += tempoExecucaoOtimizado[cont]
			fmt.Println("Determinante:", determinante)
			determinante = 0
		}

		mediaBase = somaBase / 3
		mediaOtimizado = somaOtimizado / 3

		fmt.Println("Tempos (baseline):", tempoExecucaoBase)
		fmt.Println("Média de tempo (baseline):", mediaBase, "ns")

		fmt.Println("Tempos (otimizado):", tempoExecucaoOtimizado)
		fmt.Println("Média de tempo (otimizado):", mediaOtimizado, "ns")

	}

}
