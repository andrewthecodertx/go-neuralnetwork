package main

import (
	"fmt"
	"math"
	"math/rand"
)

type NeuralNetwork struct {
	inputs       int
	hiddenLayers int
	neurons      int
	outputs      int
	weights      [][]float64
	biases       [][]float64
}

func initNetwork(inputs, hiddenLayers, neurons, outputs int) *NeuralNetwork {
	nn := NeuralNetwork{}
	nn.inputs = inputs
	nn.hiddenLayers = hiddenLayers
	nn.neurons = neurons
	nn.outputs = outputs

	nn.weights = make([][]float64, nn.inputs)
	for i := range nn.weights {
		nn.weights[i] = make([]float64, nn.hiddenLayers)
		for j := range nn.weights[i] {
			nn.weights[i][j] = rand.NormFloat64() / math.Sqrt(float64(nn.hiddenLayers))
		}
	}

	nn.biases = make([][]float64, nn.hiddenLayers)
	for i := range nn.biases {
		nn.biases[i] = make([]float64, nn.neurons)
		for j := range nn.biases[i] {
			nn.biases[i][j] = rand.Float64()
		}
	}

	return &nn
}

func (nn *NeuralNetwork) feedForward(input []float64) float64 {
	inputs := make([][]float64, 1)
	inputs[0] = make([]float64, len(input))
	copy(inputs[0], input)

<<<<<<< HEAD
	product := dotProduct(inputs, nn.weights)

	prediction := product
	for i := 0; i < nn.hiddenLayers; i++ {
		for j := 0; j < nn.neurons; j++ {
			prediction += nn.biases[i][j]
		}
=======
	for i, input := range inputs {
		hiddenLayers := make([][]float64, nn.hiddenLayerCount)

		fmt.Printf("inputs %d %g\n", i+1, inputs[i])
		for layer := 0; layer < nn.hiddenLayerCount; layer++ {
			hiddenLayers[layer] = make([]float64, nn.hiddenLayerSize)
		}

		for neuron := 0; neuron < nn.hiddenLayerSize; neuron++ {
			for inputIndex := 0; inputIndex < nn.inputSize; inputIndex++ {
				hiddenLayers[0][neuron] += input[inputIndex] * nn.weightsInputHidden[inputIndex][neuron]
			}
			hiddenLayers[0][neuron] += nn.biasHidden[0][neuron]
			hiddenLayers[0][neuron] = sigmoid(hiddenLayers[0][neuron])
		}

		for layer := 1; layer < nn.hiddenLayerCount; layer++ {
			for neuron := 0; neuron < nn.hiddenLayerSize; neuron++ {
				for prev := 0; prev < nn.hiddenLayerSize; prev++ {
					hiddenLayers[layer][neuron] += hiddenLayers[layer-1][prev] * nn.weightsHiddenOutput[layer-1][prev]
				}
				hiddenLayers[layer][neuron] += nn.biasHidden[layer][neuron]
				hiddenLayers[layer][neuron] = sigmoid(hiddenLayers[layer][neuron])
			}
		}

		output := 0.0
		for neuron := 0; neuron < nn.hiddenLayerSize; neuron++ {
			output += hiddenLayers[nn.hiddenLayerCount-1][neuron] * nn.weightsHiddenOutput[nn.hiddenLayerCount-1][neuron]
		}
		output += nn.biasOutput[0]
		output = sigmoid(output)

		output = math.Round(output*9) + 1

		predictions[i] = output
>>>>>>> 14b0a31 (semantic updates)
	}

	prediction = sigmoid(prediction)
	return prediction
}

<<<<<<< HEAD
func (nn *NeuralNetwork) train(inputs [][]float64, targets [][]float64, learnRate float64) {
	// TODO: set epochs in the man function and send it as a parameter
	for epoch := 0; epoch < 1000; epoch++ {
=======
func random() float64 {
	return 2*rand.Float64() - 1
}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func sigmoidDerivative(x float64) float64 {
	return sigmoid(x) * (1 - sigmoid(x))
}

func calculateDeviation(expected, predicted float64) float64 {
	return math.Abs(expected - predicted)
}

func backPropagation(nn *NeuralNetwork, inputs [][]float64, expectedOutputs []float64, learnRate float64, maxEpochs int, expectedError float64) int {
	epoch := 0

	for epoch < maxEpochs {
		fmt.Println("epoch", epoch+1)

		predictions := nn.feedForward(inputs)
		errors := make([]float64, len(inputs))

>>>>>>> 14b0a31 (semantic updates)
		for i := range inputs {
			// forward pass
			prediction := nn.feedForward(inputs[i])

			// loss calculation
			// TODO: this assumes a single target! need to iterate in cases where there might be more.
			loss := calculateLoss(prediction, targets[i][0])

			// backward pass
			gradients := calculateGradients(inputs[i], loss)

			// adjust weights and biases
			nn.updateWeightsAndBiases(gradients, learnRate)
		}
<<<<<<< HEAD
=======

		totalError := 0.0
		for _, err := range errors {
			totalError += math.Abs(err)
		}

		if totalError < expectedError*float64(len(inputs)) {
			fmt.Printf("training completed in %d epochs\n", epoch+1)
			return epoch + 1
		}

		fmt.Println("adjusting parameters")
		adjustParameters(nn, inputs, learnRate, errors)

		epoch++
>>>>>>> 14b0a31 (semantic updates)
	}
}

func (nn *NeuralNetwork) updateWeightsAndBiases(gradients [][]float64, learnRate float64) {
}

func calculateGradients(input []float64, loss float64) [][]float64 {
	return [][]float64{}
}

<<<<<<< HEAD
func calculateLoss(prediction, target float64) float64 {
	diff := prediction - target
	return 0.5 * diff * diff
=======
		for layerIndex := nn.hiddenLayerCount - 1; layerIndex >= 0; layerIndex-- {
			// look at each neuron in the current layer
			for neuronIndex := 0; neuronIndex < nn.hiddenLayerSize; neuronIndex++ {
				gradient := sigmoidDerivative(activation(input, nn.weightsInputHidden[neuronIndex], nn.biasOutput[neuronIndex])) * error
				// update weights between current and previous layer
				for prevNeuronIndex := 0; prevNeuronIndex < nn.hiddenLayerSize; prevNeuronIndex++ {
					nn.weightsHiddenOutput[layerIndex][neuronIndex] += learnRate * gradient * nn.weightsHiddenOutput[layerIndex-1][prevNeuronIndex]
				}
			}
		}
	}
>>>>>>> 14b0a31 (semantic updates)
}

func main() {
	var file string
	var inputCount int
	var outputCount int

	defaultFile := "data.csv"
	defaultInputCount := 11
	defaultOutputCount := 1

	fmt.Printf("data file (default: %s): ", defaultFile)
	fmt.Scanln(&file)
	if file == "" {
		file = defaultFile
	}

	fmt.Printf("inputs (default: %d): ", defaultInputCount)
	fmt.Scanln(&inputCount)
	if inputCount == 0 {
		inputCount = defaultInputCount
	}

<<<<<<< HEAD
	fmt.Printf("outputs (default: %d): ", defaultOutputCount)
	fmt.Scanln(&outputCount)
	if outputCount == 0 {
		outputCount = defaultOutputCount
=======
	epochs := 1000
	expectedError := 0.01
	learnRate := 0.1

	fmt.Println("training started")
	fmt.Println("records:", len(inputs))

	epochsCompleted := backPropagation(nn, inputs, expectedOutputs, learnRate, epochs, expectedError)

	if epochsCompleted == epochs {
		fmt.Println("training complete. max epochs reached")
	} else {
		fmt.Printf("traning finished in %d epochs\n", epochsCompleted)
>>>>>>> 14b0a31 (semantic updates)
	}

	inputs, targets := loadCSV(file, inputCount, outputCount)

	nn := initNetwork(inputCount, 4, 4, outputCount)

	// TODO: set learn rate in the main function and send it as a parameter
	nn.train(inputs, targets, 0.1)
}
