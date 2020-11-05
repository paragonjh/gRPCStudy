#include <iostream>
#include <fstream>
#include <string.h>
#include <time.h>
#include <pthread.h>
#include "libVNetClient.h"
using namespace std;

char imgPath[4][100] = {
    "/Users/jihoon/study/gRPCStudy/client_streaming_api/pcbook/tmp/test1.jpg",
    "/Users/jihoon/study/gRPCStudy/client_streaming_api/pcbook/tmp/test2.jpg",
    "/Users/jihoon/study/gRPCStudy/client_streaming_api/pcbook/tmp/test3.jpg",
    "/Users/jihoon/study/gRPCStudy/client_streaming_api/pcbook/tmp/test4.jpg",
};

void *sendImage(void *imgPath)
{
    char *imageBuf = nullptr;
    ifstream fin;
    fin.open((char *)imgPath, ios::binary);

    fin.seekg(0, ios::end);
    int imgLen = fin.tellg();
    cout << "image size: " <<  imgLen << endl;
    fin.seekg(0, ios::beg);
	imageBuf = (char*)malloc(imgLen);
    fin.read(imageBuf, imgLen);
    UploadImage(imageBuf, imgLen);

    if(imageBuf != nullptr)
        free(imageBuf);
    return 0;
}

int main()
{
    clock_t start, end;
    double result;
    int sum = 0;
    pthread_t pThread[4];
    int threadId[4] = {0,};
    int status;

    cout << "main process start" << endl;

    start = clock();
    for(int i = 0; i < 4; i++)
    {
        threadId[i] = pthread_create(&pThread[i], NULL, sendImage, (void *)imgPath[i]);
        if (threadId[i] < 0)
        {
            perror("thread create error : ");
            exit(0);
        }
    }

    for(int i = 0; i < 4; i++)
    {
        pthread_join(pThread[i], (void **)&status);
    }

    end = clock();
    result = (double)(end - start)/CLOCKS_PER_SEC;
    cout << "main process end, Total duration:"<< result << endl;
    return 0;
}
