import QtQuick
import QtGrpc
import QtQuick.Controls
import QtQuick.Layouts

import goqtgrpc
import locationsvc.v1

ApplicationWindow {
    id: root
    property int port: initialPort || 0

    property positionRequest posReq
    property string responseText: ""

    // Factory function to create a gRPC client with a custom hostUri
    function createGrpcClient(hostUri: string): LocationServiceClient {
        // Create the GrpcHttp2Channel
        var grpcChannel = Qt.createQmlObject(`
            import QtGrpc
            GrpcHttp2Channel {
                hostUri: "${hostUri}"
            }
        `, root, "dynamicGrpcChannel");

        // Create the LocationServiceClient and explicitly set its channel
        var grpcClient = Qt.createQmlObject(`
            import locationsvc.v1
            LocationServiceClient {
                id: client
            }
        `, root, "dynamicGrpcClient");

        // Explicitly set the channel for the client
        grpcClient.channel = grpcChannel.channel;

        return grpcClient;
    }


    function requestPosition(deviceId: string): void {
        root.responseText = "";
        root.posReq.deviceId = deviceId;

        // Create a new gRPC client with the desired hostUri
        var grpcClient = createGrpcClient(`http://localhost:${port}`);

        // Make the gRPC call
        grpcClient.Position(root.posReq, finishCallback, errorCallback, grpcCallOptions);
    }

    function finishCallback(response): void {
        console.log(response);
        root.responseText = JSON.stringify(response);
    }

    function errorCallback(error): void {
        // error is received as a JavaScript object, but it is a QGrpcStatus instance
        console.log(
            `Error callback executed. Error message: "${error.message}" Code: ${error.code}`
        );
        root.responseText = error.message;
    }

    minimumWidth: rootLayout.implicitWidth + rootLayout.anchors.margins * 2
    minimumHeight: rootLayout.implicitHeight + rootLayout.anchors.margins * 2

    visible: true
    title: qsTr("Demo Go GRPC Example")
    font.pointSize: 16

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
            onClicked: root.requestPosition(questionInput.text)
            Layout.alignment: Qt.AlignCenter
            leftPadding: 16
            rightPadding: 16
            text: qsTr("Request")
        }

        TextField {
            id: responseText
            Layout.alignment: Qt.AlignCenter
            Layout.fillWidth: true
            leftPadding: 10
            rightPadding: 10
            readOnly: true
            text: root.responseText
            wrapMode: Text.Wrap
        }
    }
}
