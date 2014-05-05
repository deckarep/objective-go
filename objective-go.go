package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework Foundation
#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>

NSData *data;

void
Screen(void) {
	CGImageRef image = CGDisplayCreateImage(kCGDirectMainDisplay);
	CFMutableDataRef mutableData = CFDataCreateMutable(NULL, 0);
	CGImageDestinationRef idst = CGImageDestinationCreateWithData(
		mutableData, kUTTypeJPEG, 1, NULL
	);

	NSInteger exif             =    1;
	CGFloat compressionQuality = 0.70;

	NSDictionary *props = [
		[NSDictionary alloc]
		initWithObjectsAndKeys:[NSNumber numberWithFloat:compressionQuality],
		kCGImageDestinationLossyCompressionQuality,
		[NSNumber numberWithInteger:exif],
		kCGImagePropertyOrientation, nil
	];

	CGImageDestinationAddImage(idst, image, (CFDictionaryRef)props);
	CGImageDestinationFinalize(idst);
	data = [NSData dataWithData:(NSData *)mutableData];
	[props release];
	CFRelease(idst);
	CFRelease(mutableData);
	CGImageRelease(image);
}

int
Length() {
	return [data length];
}

const void*
Data() {
	return [data bytes];
}

void
Free () {
	[data release];
}
*/
import "C"
import "fmt"
import "unsafe"
import "time"
import "io/ioutil"

func main() {
	start := time.Now()
	c := 100
	fmt.Println("started at", start)
	for i := 0; i <= c; i++ {
		C.Screen()
		ioutil.WriteFile(
			"/tmp/s.jpg",
			(*[1 << 30]byte)(unsafe.Pointer(C.Data()))[0:C.Length()],
			0644,
		)
		C.Free()
	}
	delta := time.Now().Sub(start)
	fmt.Println("done in", delta, c*1e9/int(delta))
}
