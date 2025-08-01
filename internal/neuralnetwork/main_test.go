package neuralnetwork_test

import (
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"testing"

	"go-neuralnetwork/internal/data"
	"go-neuralnetwork/internal/neuralnetwork"
)

func TestInitNetwork(t *testing.T) {
	inputs := 2
	hiddenLayers := []int{3, 2}
	outputs := 1
	hiddenActivations := []string{"relu", "tanh"}
	outputActivation := "linear"

	nn := neuralnetwork.InitNetwork(inputs, hiddenLayers, outputs, hiddenActivations, outputActivation)

	if nn.NumInputs != inputs {
		t.Errorf("Expected NumInputs to be %d, got %d", inputs, nn.NumInputs)
	}
	if !reflect.DeepEqual(nn.HiddenLayers, hiddenLayers) {
		t.Errorf("Expected HiddenLayers to be %v, got %v", hiddenLayers, nn.HiddenLayers)
	}
	if nn.NumOutputs != outputs {
		t.Errorf("Expected NumOutputs to be %d, got %d", outputs, nn.NumOutputs)
	}
	if !reflect.DeepEqual(nn.HiddenActivations, hiddenActivations) {
		t.Errorf("Expected HiddenActivations to be %v, got %v", hiddenActivations, nn.HiddenActivations)
	}
	if nn.OutputActivation != outputActivation {
		t.Errorf("Expected OutputActivation to be %s, got %s", outputActivation, nn.OutputActivation)
	}

	if len(nn.HiddenWeights) != len(hiddenLayers) {
		t.Errorf("HiddenWeights dimensions mismatch")
	}
	if len(nn.OutputWeights) != outputs {
		t.Errorf("OutputWeights dimensions mismatch")
	}
	if len(nn.HiddenBiases) != len(hiddenLayers) {
		t.Errorf("HiddenBiases dimensions mismatch")
	}
	if len(nn.OutputBiases) != outputs {
		t.Errorf("OutputBiases dimensions mismatch")
	}
}

func TestFeedForward(t *testing.T) {
	nn := &neuralnetwork.NeuralNetwork{
		NumInputs:         2,
		HiddenLayers:      []int{2, 2},
		NumOutputs:        1,
		HiddenWeights:     [][][]float64{{{0.1, 0.2}, {0.3, 0.4}}, {{0.5, 0.6}, {0.7, 0.8}}},
		OutputWeights:     [][]float64{{0.9, 1.0}},
		HiddenBiases:      [][]float64{{0.0, 0.0}, {0.0, 0.0}},
		OutputBiases:      []float64{0.0},
		HiddenActivations: []string{"relu", "relu"},
		OutputActivation:  "linear",
	}
	nn.SetActivationFunctions()

	inputs := []float64{1.0, 1.0}
	hiddenOutputs, finalOutputs := nn.FeedForward(inputs)

	expectedHidden1 := []float64{0.3, 0.7}
	expectedHidden2 := []float64{0.57, 0.77}
	expectedFinal := []float64{1.283}

	for i := range expectedHidden1 {
		if math.Abs(hiddenOutputs[0][i]-expectedHidden1[i]) > 1e-9 {
			t.Errorf("Hidden output 1 mismatch at index %d: Expected %f, got %f", i, expectedHidden1[i], hiddenOutputs[0][i])
		}
	}
	for i := range expectedHidden2 {
		if math.Abs(hiddenOutputs[1][i]-expectedHidden2[i]) > 1e-9 {
			t.Errorf("Hidden output 2 mismatch at index %d: Expected %f, got %f", i, expectedHidden2[i], hiddenOutputs[1][i])
		}
	}
	if math.Abs(finalOutputs[0]-expectedFinal[0]) > 1e-9 {
		t.Errorf("Final output mismatch: Expected %f, got %f", expectedFinal[0], finalOutputs[0])
	}
}

func TestSaveAndLoadModel(t *testing.T) {
	originalNN := neuralnetwork.InitNetwork(2, []int{2, 2}, 1, []string{"relu", "tanh"}, "linear")
	originalMD := &data.ModelData{
		NN:         originalNN,
		TargetMins: []float64{1.0},
		TargetMaxs: []float64{10.0},
		InputMins:  []float64{0.0, 0.0},
		InputMaxs:  []float64{1.0, 1.0},
	}

	tmpfile, err := ioutil.TempFile("", "model-*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	filePath := tmpfile.Name()
	tmpfile.Close()
	defer os.Remove(filePath)

	err = originalMD.SaveModel(filePath)
	if err != nil {
		t.Fatalf("Failed to save model: %v", err)
	}

	loadedMD, err := data.LoadModel(filePath)
	if err != nil {
		t.Fatalf("Failed to load model: %v", err)
	}
	loadedMD.NN.SetActivationFunctions()

	if !reflect.DeepEqual(originalMD.NN.NumInputs, loadedMD.NN.NumInputs) ||
		!reflect.DeepEqual(originalMD.NN.HiddenLayers, loadedMD.NN.HiddenLayers) ||
		!reflect.DeepEqual(originalMD.NN.NumOutputs, loadedMD.NN.NumOutputs) ||
		!reflect.DeepEqual(originalMD.NN.HiddenActivations, loadedMD.NN.HiddenActivations) ||
		!reflect.DeepEqual(originalMD.NN.OutputActivation, loadedMD.NN.OutputActivation) ||
		!reflect.DeepEqual(originalMD.NN.HiddenWeights, loadedMD.NN.HiddenWeights) ||
		!reflect.DeepEqual(originalMD.NN.OutputWeights, loadedMD.NN.OutputWeights) ||
		!reflect.DeepEqual(originalMD.NN.HiddenBiases, loadedMD.NN.HiddenBiases) ||
		!reflect.DeepEqual(originalMD.NN.OutputBiases, loadedMD.NN.OutputBiases) {
		t.Errorf("Loaded model does not match original model")
	}
}
