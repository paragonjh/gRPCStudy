#include <iostream>
#include <fstream>
#include <thread>
#include <string.h>
#include <time.h>

#include "libVNetClient.h"
using namespace std;

void sendImage1()
{
    char *imageBuf = nullptr;
    ifstream fin;
    fin.open("/Users/jihoon/study/gRPCStudy/client_streaming_api/pcbook/tmp/test1.jpg", ios::binary);

    fin.seekg(0, ios::end);
    int imgLen = fin.tellg();
    cout << "image size: " <<  imgLen << endl;
    fin.seekg(0, ios::beg);
	imageBuf = (char*)malloc(imgLen);
    fin.read(imageBuf, imgLen);
    UploadImage(imageBuf, imgLen);

    if(imageBuf != nullptr)
        free(imageBuf);
}

void sendImage2()
{
    char *imageBuf = nullptr;
    ifstream fin;
    fin.open("/Users/jihoon/study/gRPCStudy/client_streaming_api/pcbook/tmp/test2.jpg", ios::binary);

    fin.seekg(0, ios::end);
    int imgLen = fin.tellg();
    cout << "image size: " <<  imgLen << endl;
    fin.seekg(0, ios::beg);
	imageBuf = (char*)malloc(imgLen);
    fin.read(imageBuf, imgLen);
    UploadImage(imageBuf, imgLen);

    if(imageBuf != nullptr)
        free(imageBuf);
}

void sendImage3()
{
    char *imageBuf = nullptr;
    ifstream fin;
    fin.open("/Users/jihoon/study/gRPCStudy/client_streaming_api/pcbook/tmp/test3.jpg", ios::binary);

    fin.seekg(0, ios::end);
    int imgLen = fin.tellg();
    cout << "image size: " <<  imgLen << endl;
    fin.seekg(0, ios::beg);
	imageBuf = (char*)malloc(imgLen);
    fin.read(imageBuf, imgLen);
    UploadImage(imageBuf, imgLen);

    if(imageBuf != nullptr)
        free(imageBuf);
}

void sendImage4()
{
    char *imageBuf = nullptr;
    ifstream fin;
    fin.open("/Users/jihoon/study/gRPCStudy/client_streaming_api/pcbook/tmp/test4.jpg", ios::binary);

    fin.seekg(0, ios::end);
    int imgLen = fin.tellg();
    cout << "image size: " <<  imgLen << endl;
    fin.seekg(0, ios::beg);
	imageBuf = (char*)malloc(imgLen);
    fin.read(imageBuf, imgLen);
    UploadImage(imageBuf, imgLen);

    if(imageBuf != nullptr)
        free(imageBuf);
}

int main()
{
    clock_t start, end;
    double result;
    int i, j;
    int sum = 0;
    cout << "main process start" << endl;

    start = clock();
    thread _t1(sendImage1);
    thread _t2(sendImage2);
    thread _t3(sendImage3);
    thread _t4(sendImage4);

    _t1.join();
    _t2.join();
    _t3.join();
    _t4.join();
    end = clock();
    result = (double)(end - start)/CLOCKS_PER_SEC;
    cout << "main process end, Total duration:"<< result << endl;
    return 0;
}
