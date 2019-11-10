import { Empty } from 'google-protobuf/google/protobuf/empty_pb';
import * as MessageService from "~/proto/MessageService_pb_service"
import * as MessageService_pb from "~/proto/MessageService_pb";

type GetMessageStreamCallBack = {
}

export default class GrpcClient {
    private messageService: MessageService.MessageServiceClient

    public constructor () {
        this.messageService = new MessageService.MessageServiceClient("http://0.0.0.0:8080")
    }

    public getMessageStream (cb: (messageList: MessageService_pb.Message.AsObject) => void) {
        const messageStream = this.messageService.getMessageStream(new Empty())
        messageStream.on('data', message => {
            cb(message.toObject())
        })
    }

    public postMessage (message: MessageService_pb.Message) {
        this.messageService.postMessage(message, (err, res) => {
            console.log(err, res)
        })
    }

}
