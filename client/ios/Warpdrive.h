//
//  Warpdrive.h
//  Warpdrive
//
//  Created by Ali Najafizadeh on 2016-12-05.
//  Copyright © 2016 Ali Najafizadeh. All rights reserved.
//

#import <Foundation/Foundation.h>

#import "RCTBridgeModule.h"

@interface Warpdrive : NSObject

+ (instancetype)shared;

- (NSURL *)sourceBundle;
- (void) call;

@end
