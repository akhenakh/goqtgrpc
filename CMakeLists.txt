set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
cmake_minimum_required(VERSION 3.16)
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

qt_policy(SET QTP0001 NEW)



# Create the executable with source files
qt_add_executable(MyApp
    main.cpp
)

qt_add_qml_module(MyApp
    URI goqtgrpc
    VERSION 1.0
    QML_FILES
        "Main.qml"
    DEPENDENCIES
        goqtgrpc_plugin
)

target_link_libraries(MyApp
    PRIVATE
        Qt6::Protobuf
        Qt6::ProtobufWellKnownTypes
        Qt6::Core
        Qt6::Quick
        Qt6::Protobuf
        Qt6::Grpc
        Qt6::GrpcQuick
        ${CMAKE_CURRENT_SOURCE_DIR}/cposlib/cposlib.a
        goqtgrpc_plugin
)

target_include_directories(MyApp PRIVATE cposlib)
