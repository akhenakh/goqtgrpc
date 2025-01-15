import QtQuick
import QtGrpc
import QtQuick.Controls
import QtQuick.Layouts

import goqtgrpc
import locationsvc.v1

ApplicationWindow {
    id: root

    property positionRequest posReq
    property string errorText: ""
    property int errorCode: 0

    function requestAnswer(id: string): void {
        root.errorText = "";
        root.posReq.id = id;
        grpcClient.Position(root.posReq, finishCallback, errorCallback, grpcCallOptions);
    }

    function finishCallback(response: positionResponse): void {
        console.log(positionResponse);
    }

    function errorCallback(error): void {
        // error is received as a JavaScript object, but it is a QGrpcStatus instance
        console.log(
            `Error callback executed. Error message: "${error.message}" Code: ${error.code}`
        );
        root.errorText = error.message;
        root.errorCode = error.code;
    }

    minimumWidth: rootLayout.implicitWidth + rootLayout.anchors.margins * 2
    minimumHeight: rootLayout.implicitHeight + rootLayout.anchors.margins * 2

    visible: true
    title: qsTr("Demo Go GRPC Example")
    font.pointSize: 18

    GrpcHttp2Channel {
        id: grpcChannel
        hostUri: "http://localhost:" + ConfigModule.port
        // Optionally, you can specify custom channel options here
        // options: GrpcChannelOptions {}
    }
    LocationServiceClient {
        id: grpcClient
        channel: grpcChannel.channel
    }

    GrpcCallOptions {
        id: grpcCallOptions
        deadlineTimeout: 6000
    }

    ColumnLayout {
        id: rootLayout
        anchors.margins: 10
        anchors.fill: parent
        spacing: 12

        TextField {
            id: questionInput
            Layout.alignment: Qt.AlignCenter
            Layout.minimumWidth: 300
            leftPadding: 10
            rightPadding: 10
            placeholderText: qsTr("ID...")
        }

        Button {
            onClicked: root.requestAnswer(questionInput.text)
            Layout.alignment: Qt.AlignCenter
            leftPadding: 16
            rightPadding: 16
            text: qsTr("Request")
        }


    }
}
