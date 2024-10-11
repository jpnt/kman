package pkg

type State interface {
	Enter()
	Execute()
	Exit()
}
