#!/usr/bin/perl
use warnings;
use strict;

die "Usage: $0 <width> <height> <density>" unless (scalar(@ARGV) == 3);

my ($width, $height, $density) = @ARGV;

# 密度を0から100の範囲に制限
$density = 0 if $density < 0;
$density = 100 if $density > 100;

print "$width.ox\n";

for (my $i = 0; $i < $height; $i++) {
    for (my $j = 0; $j < $width; $j++) {
        if (int(rand(100)) < $density) {
            print "o";
        }
        else {
            print ".";
        }
    }
    print "\n";
}
