package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		stageCh := make(Bi)

		go func(inCh Bi, outCh Out) {
			defer close(inCh)

			for {
				select {
				case <-done:
					return
				case result, ok := <-outCh:
					if !ok {
						return
					}
					inCh <- result
				}
			}
		}(stageCh, out)

		out = stage(stageCh)
	}

	return out
}
