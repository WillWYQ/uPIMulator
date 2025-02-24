#!/bin/bash

# Define the values to replace
VALUES=(1 2 4 8 16 32 64 128)

# Iterate over each value
for NUM_DPUS in "${VALUES[@]}"; do
    BIN_DIR="bin${NUM_DPUS}"
    ZIP_FILE="RED_${NUM_DPUS}.zip"

    echo "Running uPIMulator RED case with --num_dpus_per_rank=${NUM_DPUS}"

    # Ensure a clean working directory
    rm -rf "${BIN_DIR}"
    mkdir -p "${BIN_DIR}"

    # Run the command with the current value
    ./build/uPIMulator \
        --verbose 2 \
        --root_dirpath /uPIMulator/golang_vm/uPIMulator/ \
        --bin_dirpath /uPIMulator/golang_vm/uPIMulator/"${BIN_DIR}"/ \
        --benchmark RED \
        --logic_frequency 1600 \
        --memory_frequency 3200 \
        --num_channels 8 \
        --num_ranks_per_channel 2  \
        --num_dpus_per_rank "${NUM_DPUS}" \
        --num_tasklets 16 \
        --data_prep_params 2097152 \
        --t_rcd 32 \
        --t_ras 78 \
        --t_rp 32 \
        --t_cl 32 \
        --t_bl 8

    # Zip the correct bin directory
    echo "Creating archive: ${ZIP_FILE}"
    zip -r "${ZIP_FILE}" "${BIN_DIR}"

    echo "Completed iteration for --num_dpus_per_rank=${NUM_DPUS}"
done

echo "All runs completed!"