//
//  Warpify.m
//  Warpdrive
//
//  Created by Ali Najafizadeh on 2016-12-05.
//  Copyright © 2016 Ali Najafizadeh. All rights reserved.
//

#import "RCTBundleURLProvider.h"

#import "Warpify.framework/Headers/Warpify.h"

#import "Warpify.h"
#import "ReloadTaskWrapper.h"

static Warpify *sharedInstance;

@implementation Warpify {
  // we are creating this variable to make sure it never
  // garbage collected
  ReloadTaskWrapper* _reloadTask;
}

// this method returns the document path based on whether groupName given or not
+ (NSString*)documentPathWithGroupName:(NSString*)groupName {
  if (groupName == nil) {
    NSArray *paths = NSSearchPathForDirectoriesInDomains(NSDocumentDirectory, NSUserDomainMask, YES);
    return [paths objectAtIndex:0];
  } else {
    NSURL* pathURL = [[NSFileManager defaultManager] containerURLForSecurityApplicationGroupIdentifier:groupName];
    return [pathURL path];
  }
}

+ (instancetype)createWithDefaultCycle:(NSString*)defaultCycle forceUpdate:(BOOL)forceUpdate groupName:(NSString*)groupName {
  static dispatch_once_t once_token;
  
  // we are going to call this method once so subsiqent call to `shared` and `createWithDefaultCycle`
  // returns the same instance
  dispatch_once(&once_token, ^{
    sharedInstance = [Warpify new];
    
    // need to update the internal go variable
    NSString* bundleVersion = [[NSBundle mainBundle] objectForInfoDictionaryKey:@"CFBundleShortVersionString"];
    NSString* bundlePath = [[NSBundle mainBundle] bundlePath];
    NSString* documentPath = [Warpify documentPathWithGroupName:groupName];
    NSString* platform = @"ios";

    // set the reload path here
    sharedInstance->_reloadTask = [ReloadTaskWrapper new];
    [sharedInstance->_reloadTask setBlock:^(NSString* path) {
      [sharedInstance reloadFromPath:path];
    }];
    GoWarpifySetReload((GoWarpifyTask*)sharedInstance->_reloadTask);
    
    // Setup the basic requirements
    GoWarpifySetup(bundleVersion, bundlePath, documentPath, platform, defaultCycle, forceUpdate);
  });
  
  return sharedInstance;
}

+ (instancetype)createWithDefaultCycle:(NSString*)defaultCycle forceUpdate:(BOOL)forceUpdate {
  return [Warpify createWithDefaultCycle:defaultCycle forceUpdate:forceUpdate groupName:nil];
}

+ (instancetype)shared {
  return [Warpify createWithDefaultCycle:@"prod" forceUpdate:false];
}

- (NSURL *)sourceBundle {
  return [[RCTBundleURLProvider sharedSettings] jsBundleURLForBundleRoot:@"index.ios" fallbackResource:nil];
}

- (void) reloadFromPath:(NSString*)path {
  NSLog(@"reload...");
  NSLog(path);
}

@end
