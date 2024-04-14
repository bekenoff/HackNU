package main

import (
	"math"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func speech() {
	// Create a new WAV file
	outFile, err := os.Create("c.wav")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Define the WAV file properties
	enc := wav.NewEncoder(outFile, 8000, 16, 1, 1)

	// Sample rate and duration
	sampleRate := 8000
	duration := 2 // seconds

	// Create a buffer
	buf := &audio.IntBuffer{Data: make([]int, sampleRate*duration), Format: &audio.Format{SampleRate: sampleRate, NumChannels: 1}}

	// Generate a 440 Hz sine wave (tone A4)
	freq := 440.0 // Frequency in Hz
	for i := range buf.Data {
		// Generate the samples
		buf.Data[i] = int(32767 * math.Sin(2*math.Pi*freq*float64(i)/float64(sampleRate)))
	}

	// Write the buffer to the WAV file
	if err := enc.Write(buf); err != nil {
		panic(err)
	}

	// Close the encoder which will finalize the WAV file
	if err := enc.Close(); err != nil {
		panic(err)
	}

	println("WAV file has been successfully created as 'output.wav'")
}
