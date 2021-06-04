<?php

function changeDateFormat(array $dates): array
{
    $returned = [];
    foreach ($dates as $date) {
        // YYYY/MM/DD DD/MM/YYYY MM-DD-YYYY
        $newDate = null;

        if (preg_match("/^[0-9]{4}\/(0[1-9]|1[0-2])\/(0[1-9]|[1-2][0-9]|3[0-1])$/", $date)) {
            $newDate = DateTime::createFromFormat('Y/m/d', $date);
        } elseif (preg_match("/^(0[1-9]|[1-2][0-9]|3[0-1])\/(0[1-9]|1[0-2])\/[0-9]{4}$/", $date)) {
            $newDate = DateTime::createFromFormat('d/m/Y', $date);
        } elseif (preg_match("/^(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])-[0-9]{4}$/", $date)) {
            $newDate = DateTime::createFromFormat('m-d-Y', $date);
        }
        if ($newDate)
            $returned[] = $newDate->format('Ymd');

    }
    return $returned;
}

$dates = changeDateFormat(["2010/03/30", "15/12/2016", "11-15-2012", "20130720"]);
foreach ($dates as $date) {
    echo $date . "\n";
}


