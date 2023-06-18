#include <stdio.h>
#include <string.h>
#include "minioc.h"

int main() {
    // Parameters for PresignedGetObject function
    const char *minioEndpoint = "127.0.0.1:9000";
    const char *minioAccessKey = "GnFz1opVMB2awtzfp66y";
    const char *minioSecretKey = "0SV6TCcUg2CqeMKe304BsURNnH0KeTjcmymAZAS3";
    const char *bucketName = "ttt";
    const char *objectName = "Sample-MP4-Video-File-Download.mp4";
    int expiry = 85400;

    // Call the PresignedGetObject function
    struct PresignedGetObject_return result = PresignedGetObject((void*)minioEndpoint, strlen(minioEndpoint),
                                                                  (void*)minioAccessKey, strlen(minioAccessKey),
                                                                  (void*)minioSecretKey, strlen(minioSecretKey),
                                                                  (void*)bucketName, strlen(bucketName),
                                                                  (void*)objectName, strlen(objectName),
                                                                  expiry);

    // Check the result
    if (result.r1 > 0) {
        printf("Presigned URL: %s\n", (char*)result.r0);
    } else {
        printf("Error generating presigned URL %d\n", result.r1);
    }

    return 0;
}
