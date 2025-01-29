import QtQuick
import QtGrpc
import QtQuick.Controls
import QtQuick.Layouts
import QtLocation
import QtPositioning

import goqtgrpc
import locationsvc.v1

ApplicationWindow {
    id: root
    property int port: initialPort || 0

    property positionRequest posReq
    property string responseText: ""
    property LocationServiceClient grpcClient
    property var receivedPosition: QtPositioning.coordinate() // Holds the received position

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

    function requestStreamPosition(deviceId: string): void {
        root.responseText = "";
        root.posReq.deviceId = deviceId;

        // Create a new gRPC client with the desired hostUri only if the root one is nil
        if (!root.grpcClient) {
            root.grpcClient = createGrpcClient(`http://localhost:${port}`);
        }

        // Define the callbacks for the stream
        var streamCallbacks = {
            // Callback for handling each position response
            positionReceived: function(response) {
                console.log("Stream response received:", response);
                root.responseText = JSON.stringify(response)
                receivedPosition.latitude = response.latitude;
                receivedPosition.longitude = response.longitude;
                map.center = receivedPosition; // Update the map center
                marker.coordinate = receivedPosition;
                console.log("New position:", receivedPosition.latitude, receivedPosition.longitude, receivedPosition.isValid)
            },

            // Callback for handling stream errors
            errorOccurred: function(error) {
                console.log(
                    `Stream error occurred. Error message: "${error.message}" Code: ${error.code}`
                );
                root.responseText += "Stream error: " + error.message + "\n";
            },

            // Callback for handling stream completion
            finished: function() {
                console.log("Stream finished.");
                root.responseText += "Stream finished.\n";
            }
        };

        // Make the gRPC stream call
        grpcClient.StreamPosition(root.posReq, streamCallbacks.positionReceived, streamCallbacks.errorOccurred, streamCallbacks.finished, grpcCallOptions);
    }

    function requestUnaryPosition(deviceId: string): void {
        root.responseText = "";
        root.posReq.deviceId = deviceId;

        // Create a new gRPC client with the desired hostUri only if the root one is nil
        if (!root.grpcClient) {
            root.grpcClient = createGrpcClient(`http://localhost:${port}`);
        }

        // Define the callbacks for the stream
        var unaryCallbacks = {
            // Callback for handling unary errors
            errorOccurred: function(error) {
                console.log(
                    `Stream error occurred. Error message: "${error.message}" Code: ${error.code}`
                );
                root.responseText += "error: " + error.message + "\n";
            },

            // Callback for handling unary completion
            finished: function(response) {
                console.log("Stream response received:", response);
                root.responseText = JSON.stringify(response);
                receivedPosition.latitude = response.latitude;
                receivedPosition.longitude = response.longitude;
                map.center = receivedPosition; // Update the map center
                marker.coordinate = receivedPosition;
                console.log("New position:", receivedPosition.latitude, receivedPosition.longitude)
            }
        };

        // Make the gRPC unary call
        grpcClient.Position(root.posReq, unaryCallbacks.finished, unaryCallbacks.errorOccurred, grpcCallOptions);
    }

    minimumWidth: rootLayout.implicitWidth + rootLayout.anchors.margins * 2
    minimumHeight: rootLayout.implicitHeight + rootLayout.anchors.margins * 2

    visible: true
    title: qsTr("Demo Go GRPC Example")
    font.pointSize: 16

    Plugin {
        id: mapPlugin
        name: "osm"
    }

    GrpcCallOptions {
        id: grpcCallOptions
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

        RowLayout {
            Layout.alignment: Qt.AlignCenter
            spacing: 12

            Button {
                onClicked: root.requestUnaryPosition(questionInput.text)
                leftPadding: 16
                rightPadding: 16
                text: qsTr("Request Unary")
            }
            Button {
                onClicked: root.requestStreamPosition(questionInput.text)
                leftPadding: 16
                rightPadding: 16
                text: qsTr("Request Stream")
            }
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

        Map {
            id: map
            property alias marker: marker

            Layout.fillWidth: true
            Layout.fillHeight: true
            Layout.minimumHeight: 500

            plugin: mapPlugin

            center: receivedPosition
            zoomLevel: 4

            onCenterChanged: {
                console.log("Map center changed to:", center.latitude, center.longitude);
            }

            MapCircle {
                center {
                    latitude: receivedPosition.latitude
                    longitude: receivedPosition.longitude
                }
                radius: 5000.0
                color: 'green'
                border.width: 3
            }


            // Add a marker for the received position
            MapQuickItem {
                id: marker
                anchorPoint.x: image.width/2
                anchorPoint.y: image.height
                coordinate: receivedPosition
                //visible: receivedPosition.isValid

                sourceItem: Image {
                    source: "qrc:/sat.png"
                    sourceSize: Qt.size(width, height)
                    antialiasing: true
                    smooth: true
                    // Remove any background color
                    layer.enabled: true
                    layer.effect: ShaderEffect {
                        property color color: "transparent"
                    }
                }

                onCoordinateChanged: {
                    console.log("Marker coordinate changed to:", coordinate.latitude, coordinate.longitude);
                }
            }

            // Enable panning and zooming with the mouse
            MouseArea {
            anchors.fill: parent
            acceptedButtons: Qt.LeftButton | Qt.RightButton

            property real lastX: 0
            property real lastY: 0

            onPressed: {
                lastX = mouse.x
                lastY = mouse.y
            }

            onPositionChanged: {
                if (pressedButtons & Qt.LeftButton) {
                    var dx = mouse.x - lastX
                    var dy = mouse.y - lastY
                    map.pan(-dx, -dy)
                    lastX = mouse.x
                    lastY = mouse.y
                }
            }

            onWheel: {
                if (wheel.angleDelta.y > 0) {
                    map.zoomLevel += 0.5
                } else {
                    map.zoomLevel -= 0.5
                }
            }
        }
        }
    }
}
