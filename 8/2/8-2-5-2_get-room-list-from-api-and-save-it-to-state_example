import axios from "axios"

class Rooms extends Component {
    constructor(props) {
        super(props)
        this.state = {
            rooms: []
        }
    }

    async componentDidMount() {
        const rooms = await axios.get('Room一覧APIのURL')
        this.setState({ rooms })
    }
}
