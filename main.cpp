#include <QtGui/QGuiApplication>
#include <QtQml/QQmlApplicationEngine>
#include <QQmlContext>
#include <QDebug>
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

    qInfo("starting Go loop" );
    // starts the go loop
    int port = Start();

    qInfo() << "Go loop started on port " << port;

    engine.rootContext()->setContextProperty("initialPort", port);
    engine.loadFromModule("goqtgrpc", "Main");

    return app.exec();
}
