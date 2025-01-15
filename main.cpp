#include <QtGui/QGuiApplication>
#include <QtQml/QQmlApplicationEngine>
#include "cposlib.h"

int main(int argc, char *argv[])
{
    QGuiApplication app(argc, argv);
    QQmlApplicationEngine engine;
    QObject::connect(
        &engine, &QQmlApplicationEngine::objectCreationFailed, &app,
        []() {
            QCoreApplication::exit(-1);
            Stop();
        }, Qt::QueuedConnection);

    // starts the go loop
    int port = Start();

    engine.loadFromModule("goqtgrpc", "Main");

    // Get the ConfigModule singleton and set the port
    QObject* rootObject = engine.rootObjects().first();
    QObject* configModule = engine.singletonInstance<QObject*>("ConfigModule", "ConfigModule");
    if (configModule) {
        configModule->setProperty("port", port);
    }


    return app.exec();
}
