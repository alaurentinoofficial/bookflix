import { Server } from '../server'
import { cyan, blue, green, bold } from 'colors'

Server.listen(Server.get('port'), () => {
	console.log(green('➜  ') + bold(cyan('SERVER:')) + " Listening in " + blue(`:${Server.get('port')}`))
})