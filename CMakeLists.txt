set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
cmake_minimum_required(VERSION 3.16...3.22)
project(GoQtgRPC)

find_package(Qt6 REQUIRED COMPONENTS
    Core
    Protobuf
    ProtobufQuick
    Grpc
    GrpcQuick
    GrpcTools
    Quick
    ProtobufWellKnownTypes
)

qt_standard_project_setup()

qt_add_protobuf(goqtgrpc_plugin
    GENERATE_PACKAGE_SUBFOLDERS
    PROTO_FILES
        api/proto/locationsvc/v1/pos.proto
    QML
)

qt_add_grpc(goqtgrpc_plugin CLIENT
    PROTO_FILES
        api/proto/locationsvc/v1/pos.proto
    QML
)

qt_add_qml_module(goqtgrpc
    URI goqtgrpc
    VERSION 1.0
    RESOURCE_PREFIX "/qt/qml"
    QML_FILES
        "Main.qml"
)


qt_add_executable(MyApp main.cpp)

target_link_libraries(MyApp
    PRIVATE
        Qt6::Protobuf
        Qt6::ProtobufWellKnownTypes
        Qt6::Core
        Qt6::Quick
        Qt6::Protobuf
        Qt6::Grpc
        Qt6::GrpcQuick
        goqtgrpc_plugin
)
